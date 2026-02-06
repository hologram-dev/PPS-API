package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type ProyectoPuestoUseCase struct{}

func (ppu *ProyectoPuestoUseCase) Create(c context.Context, proyectoPuesto domain.ProyectoPuesto) error {
	db := bootstrap.DB
	err := db.Create(&proyectoPuesto)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (ppu *ProyectoPuestoUseCase) Fetch(c context.Context) ([]domain.ProyectoPuesto, error) {
	db := bootstrap.DB
	proyectoPuestos := []domain.ProyectoPuesto{}
	err := db.Find(&proyectoPuestos)
	if err.Error != nil {
		return nil, err.Error
	}
	return proyectoPuestos, nil
}

func (ppu *ProyectoPuestoUseCase) FetchById(c context.Context, id uuid.UUID) (domain.ProyectoPuesto, error) {
	db := bootstrap.DB
	proyectoPuesto := domain.ProyectoPuesto{}
	err := db.Where("id = ?", id).First(&proyectoPuesto)
	if err.Error != nil {
		return domain.ProyectoPuesto{}, err.Error
	}
	return proyectoPuesto, nil
}

func (ppu *ProyectoPuestoUseCase) Update(c context.Context, updatedProyectoPuesto domain.ProyectoPuesto) error {
	db := bootstrap.DB
	if err := db.Model(&updatedProyectoPuesto).
		Omit("deleted_at", "created_at").
		Updates(updatedProyectoPuesto).Error; err != nil {
		return err
	}
	return nil
}

func (ppu *ProyectoPuestoUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.ProyectoPuesto{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
