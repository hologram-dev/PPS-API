package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EstudianteCarrera struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	NroLegajo int            `json:"nro_legajo"`
}

type EstudianteCarreraRepository interface {
	Create(c context.Context, estudiantecarrera EstudianteCarrera) error
	Fetch(c context.Context) ([]EstudianteCarrera, error)
	FetchById(c context.Context, id uuid.UUID) (EstudianteCarrera, error)
	Update(c context.Context, updatedestudiantecarrera EstudianteCarrera) error
	Delete(c context.Context, id uuid.UUID) error
}
