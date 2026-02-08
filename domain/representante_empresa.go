package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepresentanteEmpresa struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Nombre    string         `json:"nombre"`
	Apellido  string         `json:"apellido"`
	Correo    string         `json:"correo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type RepresentanteEmpresaRepository interface {
	Create(c context.Context, representanteEmpresa RepresentanteEmpresa) error
	Fetch(c context.Context) ([]RepresentanteEmpresa, error)
	FetchById(c context.Context, id uuid.UUID) (RepresentanteEmpresa, error)
	Update(c context.Context, updatedRepresentanteEmpresa RepresentanteEmpresa) error
	Delete(c context.Context, id uuid.UUID) error
}
