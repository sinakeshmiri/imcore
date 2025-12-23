package repository

import (
	"context"
	"time"

	"github.com/sinakeshmiri/imcore/domain"
	"gorm.io/gorm"
)

func (ApplicationModel) TableName() string { return "applications" }

type ApplicationRepository struct {
	db *gorm.DB
}

func (a ApplicationRepository) Create(c context.Context, app *domain.Application) error {
	return a.db.Create(app).Error
}

func (a ApplicationRepository) GetByID(ctx context.Context, id string) (*domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) ListOutGoing(c context.Context, id string) ([]*domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) ListInComing(c context.Context, id string) ([]*domain.Application, error) {

	var models []ApplicationModel

	err := a.db.WithContext(c).
		Table("applications a").
		Select("a.*").
		Joins("JOIN roles r ON r.rolename = a.rolename").
		Where("r.owner_username = ?", id).
		Order("a.created_at DESC").
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

	Status       string `gorm:"type:varchar(16);not null;index;column:status"`
	Reason       string `gorm:"type:varchar(640);column:reason"`
	DecisionNote string `gorm:"type:varchar(640);column:decision_note"`

	CreatedAt time.Time  `gorm:"not null;default:now();column:created_at"`
	DecidedAt *time.Time `gorm:"column:decided_at"`
	UpdatedAt time.Time  `gorm:"not null;default:now();column:updated_at"`
}

func toDomain(m ApplicationModel) (domain.Application, error) {
	status, err := domain.ParseStatus(m.Status)
	if err != nil {
		return domain.Application{}, err
	}

	return domain.Application{
		ID:                m.ID,
		Rolename:          m.Rolename,
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
