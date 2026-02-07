package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlanEstudios struct {
	ID                      uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt               time.Time      `json:"created_at"`
	UpdatedAt               time.Time      `json:"updated_at"`
	DeletedAt               gorm.DeletedAt `json:"deleted_at"`
	FechaFinPlanEstudios    time.Time      `json:"fecha_fin_plan_estudios"`
	FechaInicioPlanEstudios time.Time      `json:"fecha_inicio_plan_estudios"`
	NombrePlanEstudios      string         `json:"nombre_plan_estudios"`
}

type PlanEstudiosRepository interface {
	Create(c context.Context, planEstudios PlanEstudios) error
	Fetch(c context.Context) ([]PlanEstudios, error)
	FetchById(c context.Context, id uuid.UUID) (PlanEstudios, error)
	Update(c context.Context, updatedPlanEstudios PlanEstudios) error
	Delete(c context.Context, id uuid.UUID) error
}
