package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type MateriaUseCase struct{}

func (mu *MateriaUseCase) Create(c context.Context, materia domain.Materia) error {
	db := bootstrap.DB
	err := db.Create(&materia)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (mu *MateriaUseCase) Fetch(c context.Context) ([]domain.Materia, error) {
	db := bootstrap.DB
	materia := []domain.Materia{}
	err := db.Find(&materia)
	if err.Error != nil {
		return nil, err.Error
	}
	return materia, nil
}

func (mu *MateriaUseCase) FetchById(c context.Context, id int) (domain.Materia, error) {
	db := bootstrap.DB
	materia := domain.Materia{}
	err := db.Where("id = ?", id).First(&materia)
	if err.Error != nil {
		return domain.Materia{}, err.Error
	}
	return materia, nil
}

func (mu *MateriaUseCase) Update(c context.Context, updatedMateria domain.Materia) error {
	db := bootstrap.DB
	if err := db.Model(&updatedMateria).
		Omit("deleted_at", "created_at").
		Updates(updatedMateria).Error; err != nil {
		return err
	}
	return nil
}

func (mu *MateriaUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Materia{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
