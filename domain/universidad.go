package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Universidad struct {
	ID                   uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at"`
	NombreUniversidad    string         `json:"nombre_universidad"`
	CuitUniversidad      string         `json:"cuit_universidad"`
	DireccionUniversidad string         `json:"direccion_universidad"`
	NroTelefono          string         `json:"telefono_universidad"`
	CodigoPostal         string         `json:"codigo_postal"`
	CorreoUniversidad    string         `json:"correo_universidad"`
}

type UniversidadRepository interface {
	Fetch(c context.Context) ([]Universidad, error)
	FetchById(c context.Context, id uuid.UUID) (Universidad, error)
	Create(c context.Context, universidad Universidad) error
	Update(c context.Context, updatedUniversidad Universidad) error
	Delete(c context.Context, id uuid.UUID) error
}
