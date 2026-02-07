package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Proyecto struct {
	ID                           uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	DeletedAt                    gorm.DeletedAt `json:"deleted_at"`
	DescripcionProyecto          string         `json:"descripcion_proyecto"`
	FechaFinProyecto             time.Time      `json:"fecha_fin_proyecto"`
	FechaHoraCierrePostulaciones time.Time      `json:"fecha_hora_cierre_postulaciones"`
	FechaHoraInicioPostulaciones time.Time      `json:"fecha_hora_inicio_postulaciones"`
	FechaInicioActividades       time.Time      `json:"fecha_inicio_actividades"`
	FechaFinActividades          time.Time      `json:"fecha_fin_actividades"`
	NombreProyecto               string         `json:"nombre_proyecto"`
}

type ProyectoRepository interface {
	Create(c context.Context, proyecto Proyecto) error
	Fetch(c context.Context) ([]Proyecto, error)
	FetchById(c context.Context, id uuid.UUID) (Proyecto, error)
	Update(c context.Context, updatedProyecto Proyecto) error
	Delete(c context.Context, id uuid.UUID) error
}
