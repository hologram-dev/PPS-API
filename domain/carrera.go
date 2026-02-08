package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Carrera struct {
	ID               uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	IDExterno        string         `json:"id_externo"`
	Nombre           string         `json:"nombre"`
	Descripcion      string         `json:"descripcion"`
	FechaBajaCarrera string         `json:"fecha_baja_carrera"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type CarreraRepository interface {
	Create(c context.Context, carrera Carrera) error
	Fetch(c context.Context) ([]Carrera, error)
	FetchById(c context.Context, id uuid.UUID) (Carrera, error)
	Update(c context.Context, updatedCarrera Carrera) error
	Delete(c context.Context, id uuid.UUID) error
}
