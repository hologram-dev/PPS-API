package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProcesoSeleccionEstado struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FechaInicio time.Time      `json:"fecha_inicio"`
	FechaFin    time.Time      `json:"fecha_fin"`
	FechaCambio time.Time      `json:"fecha_cambio"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ProcesoSeleccionEstadoRepository interface {
	Create(c context.Context, procesoSeleccionEstado ProcesoSeleccionEstado) error
	Fetch(c context.Context) ([]ProcesoSeleccionEstado, error)
	FetchById(c context.Context, id uuid.UUID) (ProcesoSeleccionEstado, error)
	Update(c context.Context, updatedProcesoSeleccionEstado ProcesoSeleccionEstado) error
	Delete(c context.Context, id uuid.UUID) error
}
