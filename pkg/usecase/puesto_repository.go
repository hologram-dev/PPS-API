package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type PuestoUseCase struct{}

func (eu *PuestoUseCase) Create(c context.Context, puesto domain.Puesto) error {
	db := bootstrap.DB
	err := db.Create(&puesto)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *PuestoUseCase) Fetch(c context.Context) ([]domain.Puesto, error) {
	db := bootstrap.DB
	entity := []domain.Puesto{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *PuestoUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Puesto, error) {
	db := bootstrap.DB
	puesto := domain.Puesto{}
	err := db.Where("id = ?", id).First(&puesto)
	if err.Error != nil {
		return domain.Puesto{}, err.Error
	}
	return puesto, nil
}

func (eu *PuestoUseCase) Update(c context.Context, updatedPuesto domain.Puesto) error {
	db := bootstrap.DB
	if err := db.Model(&updatedPuesto).
		Omit("deleted_at", "created_at").
		Updates(updatedPuesto).Error; err != nil {
		return err
	}
	return nil
}

func (eu *PuestoUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Puesto{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
