package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Postulacion struct {
	ID                    uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
	CantMateriasAprobadas int            `json:"cant_materias_aprobadas"`
	CantMateriasRegulares int            `json:"cant_materias_regulares"`
	NumeroPostulacion     int            `json:"numero_postulacion"`
}

type PostulacionRepository interface {
	Fetch(c context.Context) ([]Postulacion, error)
	FetchById(c context.Context, id uuid.UUID) (Postulacion, error)
	Create(c context.Context, postulacion Postulacion) error
	Update(c context.Context, updatedPostulacion Postulacion) error
	Delete(c context.Context, id uuid.UUID) error
}
