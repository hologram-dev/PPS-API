package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Puesto struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	NombrePuesto      string         `json:"nombre_puesto"`
	DescripcionPuesto string         `json:"descripcion_puesto"`
}

type PuestoRepository interface {
	Create(c context.Context, puesto Puesto) error
	Fetch(c context.Context) ([]Puesto, error)
	FetchById(c context.Context, id uuid.UUID) (Puesto, error)
	Update(c context.Context, updatedPuesto Puesto) error
	Delete(c context.Context, id uuid.UUID) error
}
