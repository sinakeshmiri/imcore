package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sinakeshmiri/authon-core/internal/applications/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (a ApplicationModel) TableName() string { return "applications" }

type ApplicationRepository struct {
	db *gorm.DB
}

func (a *ApplicationRepository) Create(c context.Context, roleName string, username string, reason string) (domain.Application, error) {
	var ownerUsername string

	err := a.db.WithContext(c).
		Table("roles").
		Select("owner_username").
		Where("rolename = ?", roleName).
		Scan(&ownerUsername).Error

	if err != nil {
		return domain.Application{}, err
	}
	if ownerUsername == "" {
		return domain.Application{}, domain.ErrApplicationNotFound
	}
	now := time.Now()
	app := domain.Application{
		ID:                uuid.NewString(),
		Rolename:          roleName,
		ApplicantUsername: username,
		OwnerUsername:     ownerUsername,
		Status:            domain.Pending,
		CreatedAt:         now,
		Reason:            reason,
	}

	entity := fromDomain(app)
	err = a.db.WithContext(c).Create(&entity).Error
	if err != nil {
		return domain.Application{}, err
	}

	return app, nil
}

func (a *ApplicationRepository) GetByID(ctx context.Context, id string) (*domain.Application, error) {
	var app ApplicationModel
	err := a.db.WithContext(ctx).
		Table("applications").
		Where("id = ?", id).First(&app).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrApplicationNotFound
		}
		return nil, err
	}
	application, err := toDomain(app)
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (a *ApplicationRepository) ListOutGoing(c context.Context, applicantUsername string) ([]*domain.Application, error) {
	var models []ApplicationModel

	err := a.db.WithContext(c).
		Model(&ApplicationModel{}).
		Where("applicant_username = ?", applicantUsername).
		Order("created_at DESC").
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	out := make([]*domain.Application, 0, len(models))
	for _, m := range models {
		d, err := toDomain(m)
		if err != nil {
			return nil, err
		}
		dc := d
		out = append(out, &dc)
	}

	return out, nil
}

func (a *ApplicationRepository) ListInComing(c context.Context, ownerUsername string) ([]*domain.Application, error) {
	var models []ApplicationModel
	err := a.db.WithContext(c).
		Model(&ApplicationModel{}).
		Where("owner_username = ?", ownerUsername).
		Order("created_at DESC").
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	out := make([]*domain.Application, 0, len(models))
	for _, m := range models {
		d, err := toDomain(m)
		if err != nil {
			return nil, err
		}
		dc := d
		out = append(out, &dc)
	}

	return out, nil
}
func (a *ApplicationRepository) Approve(ctx context.Context, applicationID string, decisionNote *string) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		app, err := a.decide(ctx, tx, applicationID, domain.Approved, decisionNote)
		if err != nil {
			return err
		}
		return tx.Exec(`
        INSERT INTO user_roles (username, rolename)
        VALUES (?, ?)
        ON CONFLICT DO NOTHING
    `, app.ApplicantUsername, app.Rolename).Error
	})
}

func (a *ApplicationRepository) Cancel(ctx context.Context, applicationID string, decisionNote *string) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := a.decide(ctx, tx, applicationID, domain.Canceled, decisionNote)
		return err
	})
}

func (a *ApplicationRepository) Reject(ctx context.Context, applicationID string, decisionNote *string) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := a.decide(ctx, tx, applicationID, domain.Rejected, decisionNote)
		return err
	})
}

func (a *ApplicationRepository) ExistsPending(ctx context.Context, rolename string, username string) (bool, error) {
	var count int64
	err := a.db.WithContext(ctx).
		Table("applications").
		Where("rolename = ?", rolename).
		Where("applicant_username = ?", username).
		Where("status = ?", "PENDING").
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

type ApplicationModel struct {
	ID                string `gorm:"type:uuid;primaryKey;column:application_id"`
	Rolename          string `gorm:"type:varchar(64);not null;index;column:rolename"`
	ApplicantUsername string `gorm:"type:varchar(64);not null;index;column:applicant_username"`
	OwnerUsername     string `gorm:"type:varchar(64);not null;index;column:owner_username"`

	Status       string `gorm:"type:varchar(16);not null;index;column:status"`
	Reason       string `gorm:"type:varchar(640);column:reason"`
	DecisionNote string `gorm:"type:varchar(640);column:decision_note"`

	CreatedAt time.Time  `gorm:"not null;default:now();column:created_at"`
	DecidedAt *time.Time `gorm:"column:decided_at"`
}

func toDomain(m ApplicationModel) (domain.Application, error) {
	status, err := domain.ParseStatus(m.Status)
	if err != nil {
		return domain.Application{}, err
	}

	return domain.Application{
		ID:                m.ID,
		Rolename:          m.Rolename,
		OwnerUsername:     m.OwnerUsername,
		ApplicantUsername: m.ApplicantUsername,
		Status:            status,
		Reason:            m.Reason,
		DecisionNote:      m.DecisionNote,
		CreatedAt:         m.CreatedAt,
		DecidedAt:         m.DecidedAt,
	}, nil
}
func fromDomain(d domain.Application) ApplicationModel {
	return ApplicationModel{
		ID:                d.ID,
		Rolename:          d.Rolename,
		OwnerUsername:     d.OwnerUsername,
		ApplicantUsername: d.ApplicantUsername,
		Status:            d.Status.String(),
		Reason:            d.Reason,
		DecisionNote:      d.DecisionNote,
		CreatedAt:         d.CreatedAt,
		DecidedAt:         d.DecidedAt,
	}
}
func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (a *ApplicationRepository) decide(ctx context.Context, tx *gorm.DB, applicationID string, newStatus domain.ApplicationStatus, decisionNote *string) (ApplicationModel, error) {
	var app ApplicationModel
	err := tx.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("application_id = ?", applicationID).
		First(&app).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ApplicationModel{}, domain.ErrApplicationNotFound
		}
		return ApplicationModel{}, err
	}

	current, err := domain.ParseStatus(app.Status)
	if err != nil {
		return ApplicationModel{}, err
	}
	if current != domain.Pending {
		return ApplicationModel{}, domain.ErrInvalidTransition
	}

	updates := map[string]any{
		"status":     newStatus.String(),
		"decided_at": time.Now(),
	}
	if decisionNote != nil {
		updates["decision_note"] = *decisionNote
	}

	res := tx.WithContext(ctx).
		Model(&ApplicationModel{}).
		Where("application_id = ?", applicationID).
		Where("status = ?", domain.Pending.String()).
		Updates(updates)

	if res.Error != nil {
		return ApplicationModel{}, res.Error
	}
	if res.RowsAffected == 0 {
		return ApplicationModel{}, domain.ErrInvalidTransition
	}

	return app, nil
}
