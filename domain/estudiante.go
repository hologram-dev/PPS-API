package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Estudiante struct {
	ID                        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt                 time.Time      `json:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `json:"deleted_at"`
	ApellidoEstudiante        string         `json:"apellido_estudiante"`
	CorreoEstudiante          string         `json:"correo_estudiante"`
	CorreoInstitucional       string         `json:"correo_institucional"`
	CuilEstudiante            string         `json:"cuil_estudiante"`
	DniEstudiante             string         `json:"dni_estudiante"`
	FechaNacimientoEstudiante time.Time      `json:"fecha_nacimiento_estudiante"`
	NombreEstudiante          string         `json:"nombre_estudiante"`
	TipoDNI                   string         `json:"tipo_dni"`
}

type EstudianteRepository interface {
	Fetch(c context.Context) ([]Estudiante, error)
	FetchById(c context.Context, id uuid.UUID) (Estudiante, error)
	Create(c context.Context, estudiante Estudiante) error
	Update(c context.Context, updatedEstudiante Estudiante) error
	Delete(c context.Context, id uuid.UUID) error
}
