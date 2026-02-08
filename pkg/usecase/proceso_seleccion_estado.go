package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type ProcesoSeleccionEstadoUseCase struct{}

func (seu *ProcesoSeleccionEstadoUseCase) Create(c context.Context, procesoSeleccionEstado domain.ProcesoSeleccionEstado) error {
	db := bootstrap.DB
	err := db.Create(&procesoSeleccionEstado)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (seu *ProcesoSeleccionEstadoUseCase) Fetch(c context.Context) ([]domain.ProcesoSeleccionEstado, error) {
	db := bootstrap.DB
	procesoSeleccionEstado := []domain.ProcesoSeleccionEstado{}
	err := db.Find(&procesoSeleccionEstado)
	if err.Error != nil {
		return nil, err.Error
	}
	return procesoSeleccionEstado, nil
}

func (seu *ProcesoSeleccionEstadoUseCase) FetchById(c context.Context, id uuid.UUID) (domain.ProcesoSeleccionEstado, error) {
	db := bootstrap.DB
	procesoSeleccionEstado := domain.ProcesoSeleccionEstado{}
	err := db.Where("id = ?", id).First(&procesoSeleccionEstado)
	if err.Error != nil {
		return domain.ProcesoSeleccionEstado{}, err.Error
	}
	return procesoSeleccionEstado, nil
}

func (seu *ProcesoSeleccionEstadoUseCase) Update(c context.Context, updatedProcesoSeleccionEstado domain.ProcesoSeleccionEstado) error {
	db := bootstrap.DB
	if err := db.Model(&updatedProcesoSeleccionEstado).
		Omit("deleted_at", "created_at").
		Updates(updatedProcesoSeleccionEstado).Error; err != nil {
		return err
	}
	return nil
}

func (seu *ProcesoSeleccionEstadoUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.ProcesoSeleccionEstado{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
