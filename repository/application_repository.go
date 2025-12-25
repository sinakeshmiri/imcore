package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sinakeshmiri/imcore/domain"
	"gorm.io/gorm"
)

func (a ApplicationModel) TableName() string { return "applications" }

type ApplicationRepository struct {
	db *gorm.DB
}

func (a ApplicationRepository) Create(c context.Context, roleName string, username string, reason string) (domain.Application, error) {
	var role domain.Role
	err := a.db.WithContext(c).
		Table("roles").
		Select("owner_username").
		Where("rolename = ?", roleName).
		Scan(&role).Error
	if err != nil {
		return domain.Application{}, err
	}
	now := time.Now()
	app := domain.Application{
		ID:                uuid.NewString(),
		Rolename:          roleName,
		ApplicantUsername: username,
		OwnerUsername:     role.OwnerUsername,
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

func (a ApplicationRepository) GetByID(ctx context.Context, id string) (*domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) ListOutGoing(c context.Context, applicantUsername string) ([]*domain.Application, error) {
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
		d, err := toDomain(m) // if your toDomain returns (domain.Application, error)
		if err != nil {
			return nil, err
		}
		dc := d
		out = append(out, &dc)
	}

	return out, nil
}

func (a ApplicationRepository) ListInComing(c context.Context, ownerUsername string) ([]*domain.Application, error) {
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
		d, err := toDomain(m) // (domain.Application, error)
		if err != nil {
			return nil, err
		}
		dc := d
		out = append(out, &dc)
	}

	return out, nil
}
func (a ApplicationRepository) UpdateStatus(c context.Context, id string, status domain.ApplicationStatus) error {
	//TODO implement me
	panic("implement me")
}

func (a *ApplicationRepository) ExistsPending(
	ctx context.Context,
	rolename string,
	username string,
) (bool, error) {

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
