package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProyectoPuesto struct {
	ID                      uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt               time.Time      `json:"created_at"`
	UpdatedAt               time.Time      `json:"updated_at"`
	DeletedAt               gorm.DeletedAt `json:"deleted_at"`
	CantidadSuPostulaciones int            `json:"cantidad_su_postulaciones"`
	CantidadVacantes        int            `json:"cantidad_vacantes"`
	FechaBajaProyectoPuesto *time.Time     `json:"fecha_baja_proyecto_puesto"`
	HorasDedicadas          float64        `json:"horas_dedicadas"`
}

type ProyectoPuestoRepository interface {
	Fetch(c context.Context) ([]ProyectoPuesto, error)
	FetchById(c context.Context, id uuid.UUID) (ProyectoPuesto, error)
	Create(c context.Context, proyectoPuesto ProyectoPuesto) error
	Update(c context.Context, updatedProyectoPuesto ProyectoPuesto) error
	Delete(c context.Context, id uuid.UUID) error
}
