package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contrato struct {
	ID                  uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
	NumeroContrato      int            `json:"numero_contrato"`
	FechaInicioContrato time.Time      `json:"fecha_inicio_contrato"`
	FechaFinContrato    time.Time      `json:"fecha_fin_contrato"`
}

type ContratoRepository interface {
	Fetch(c context.Context) ([]Contrato, error)
	FetchById(c context.Context, id uuid.UUID) (Contrato, error)
	Create(c context.Context, contrato Contrato) error
	Update(c context.Context, updatedContrato Contrato) error
	Delete(c context.Context, id uuid.UUID) error
}
