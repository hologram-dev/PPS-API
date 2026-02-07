package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Empresa struct {
	ID                    uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
	NombreEmpresa         string         `json:"nombre_empresa"`
	CodPostalEmpresa      string         `json:"cod_postal_empresa"`
	CuitEmpresa           string         `json:"cuit_empresa"`
	DireccionEmpresa      string         `json:"direccion_empresa"`
	NumeroTelefonoEmpresa string         `json:"numero_telefono_empresa"`
}

type EmpresaRepository interface {
	Create(c context.Context, empresa Empresa) error
	Fetch(c context.Context) ([]Empresa, error)
	FetchById(c context.Context, id uuid.UUID) (Empresa, error)
	Update(c context.Context, updatedEmpresa Empresa) error
	Delete(c context.Context, id uuid.UUID) error
}
