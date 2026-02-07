package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type EstudianteUseCase struct{}

func (eu *EstudianteUseCase) Create(c context.Context, estudiante domain.Estudiante) error {
	db := bootstrap.DB
	err := db.Create(&estudiante)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *EstudianteUseCase) Fetch(c context.Context) ([]domain.Estudiante, error) {
	db := bootstrap.DB
	estudiantes := []domain.Estudiante{}
	err := db.Find(&estudiantes)
	if err.Error != nil {
		return nil, err.Error
	}
	return estudiantes, nil
}

func (eu *EstudianteUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Estudiante, error) {
	db := bootstrap.DB
	estudiante := domain.Estudiante{}
	err := db.Where("id = ?", id).First(&estudiante)
	if err.Error != nil {
		return domain.Estudiante{}, err.Error
	}
	return estudiante, nil
}

func (eu *EstudianteUseCase) Update(c context.Context, updatedEstudiante domain.Estudiante) error {
	db := bootstrap.DB
	if err := db.Model(&updatedEstudiante).
		Omit("deleted_at", "created_at").
		Updates(updatedEstudiante).Error; err != nil {
		return err
	}
	return nil
}

func (eu *EstudianteUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Estudiante{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
