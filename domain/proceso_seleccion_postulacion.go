package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProcesoSeleccionPostulacion struct {
	ID                    uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	ContadorPostulaciones int            `json:"contador_postulaciones"`
	OrdenMerito           int            `json:"orden_merito"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}

type ProcesoSeleccionPostulacionRepository interface {
	Create(c context.Context, procesoSeleccionPostulacion ProcesoSeleccionPostulacion) error
	Fetch(c context.Context) ([]ProcesoSeleccionPostulacion, error)
	FetchById(c context.Context, id uuid.UUID) (ProcesoSeleccionPostulacion, error)
	Update(c context.Context, updatedProcesoSeleccionPostulacion ProcesoSeleccionPostulacion) error
	Delete(c context.Context, id uuid.UUID) error
}
