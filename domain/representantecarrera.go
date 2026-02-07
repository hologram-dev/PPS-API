package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepresentanteCarrera struct {
	ID                           uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	DeletedAt                    gorm.DeletedAt `json:"deleted_at"`
	ApellidoRepresentanteCarrera string         `json:"apellido"`
	CorreoRepresentanteCarrera   string         `json:"correo"`
	NombreRepresentanteCarrera   string         `json:"nombre"`
}

type RepresentanteCarreraRepository interface {
	Create(c context.Context, representanteCarrera RepresentanteCarrera) error
	Fetch(c context.Context) ([]RepresentanteCarrera, error)
	FetchById(c context.Context, id uuid.UUID) (RepresentanteCarrera, error)
	Update(c context.Context, updatedRepresentanteCarrera RepresentanteCarrera) error
	Delete(c context.Context, id uuid.UUID) error
}
