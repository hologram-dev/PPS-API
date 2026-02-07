package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostulacionEstado struct {
	ID                        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FechaCambioEstado         time.Time      `json:"fecha_cambio_estado"`
	FechaFinPostulacionEstado time.Time      `json:"fecha_fin_postulacion_estado"`
	FechaInicioPostulacion    time.Time      `json:"fecha_inicio_postulacion"`
	CreatedAt                 time.Time      `json:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `json:"deleted_at"`
}

type PostulacionEstadoRepository interface {
	Create(c context.Context, postulacionEstado PostulacionEstado) error
	Fetch(c context.Context) ([]PostulacionEstado, error)
	FetchById(c context.Context, id uuid.UUID) (PostulacionEstado, error)
	Update(c context.Context, updatedPostulacionEstado PostulacionEstado) error
	Delete(c context.Context, id uuid.UUID) error
}
