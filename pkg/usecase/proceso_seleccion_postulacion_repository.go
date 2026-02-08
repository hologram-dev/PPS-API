package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type ProcesoSeleccionPostulacionUseCase struct{}

func (spu *ProcesoSeleccionPostulacionUseCase) Create(c context.Context, procesoSeleccionPostulacion domain.ProcesoSeleccionPostulacion) error {
	db := bootstrap.DB
	err := db.Create(&procesoSeleccionPostulacion)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (spu *ProcesoSeleccionPostulacionUseCase) Fetch(c context.Context) ([]domain.ProcesoSeleccionPostulacion, error) {
	db := bootstrap.DB
	procesoSeleccionPostulacion := []domain.ProcesoSeleccionPostulacion{}
	err := db.Find(&procesoSeleccionPostulacion)
	if err.Error != nil {
		return nil, err.Error
	}
	return procesoSeleccionPostulacion, nil
}

func (spu *ProcesoSeleccionPostulacionUseCase) FetchById(c context.Context, id uuid.UUID) (domain.ProcesoSeleccionPostulacion, error) {
	db := bootstrap.DB
	procesoSeleccionPostulacion := domain.ProcesoSeleccionPostulacion{}
	err := db.Where("id = ?", id).First(&procesoSeleccionPostulacion)
	if err.Error != nil {
		return domain.ProcesoSeleccionPostulacion{}, err.Error
	}
	return procesoSeleccionPostulacion, nil
}

func (spu *ProcesoSeleccionPostulacionUseCase) Update(c context.Context, updatedProcesoSeleccionPostulacion domain.ProcesoSeleccionPostulacion) error {
	db := bootstrap.DB
	if err := db.Model(&updatedProcesoSeleccionPostulacion).
		Omit("deleted_at", "created_at").
		Updates(updatedProcesoSeleccionPostulacion).Error; err != nil {
		return err
	}
	return nil
}

func (spu *ProcesoSeleccionPostulacionUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.ProcesoSeleccionPostulacion{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
