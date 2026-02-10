package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Materia struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
}

type MateriaRepository interface {
	Create(c context.Context, materia Materia) error
	Fetch(c context.Context) ([]Materia, error)
	FetchById(c context.Context, id int) (Materia, error)
	Update(c context.Context, updatedMateria Materia) error
	Delete(c context.Context, id int) error
}
