package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type EstudianteCarreraUseCase struct{}

func (ecu *EstudianteCarreraUseCase) Create(c context.Context, estudiantecarrera domain.EstudianteCarrera) error {
	db := bootstrap.DB
	err := db.Create(&estudiantecarrera)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (ecu *EstudianteCarreraUseCase) Fetch(c context.Context) ([]domain.EstudianteCarrera, error) {
	db := bootstrap.DB
	estudiantecarrera := []domain.EstudianteCarrera{}
	err := db.Find(&estudiantecarrera)
	if err.Error != nil {
		return nil, err.Error
	}
	return estudiantecarrera, nil
}

func (ecu *EstudianteCarreraUseCase) FetchById(c context.Context, id uuid.UUID) (domain.EstudianteCarrera, error) {
	db := bootstrap.DB
	estudiantecarrera := domain.EstudianteCarrera{}
	err := db.Where("id = ?", id).First(&estudiantecarrera)
	if err.Error != nil {
		return domain.EstudianteCarrera{}, err.Error
	}
	return estudiantecarrera, nil
}

func (ecu *EstudianteCarreraUseCase) Update(c context.Context, updateEstudianteCarrera domain.EstudianteCarrera) error {
	db := bootstrap.DB
	if err := db.Model(&updateEstudianteCarrera).
		Omit("deleted_at", "created_at").
		Updates(updateEstudianteCarrera).Error; err != nil {
		return err
	}
	return nil
}

func (ecu *EstudianteCarreraUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.EstudianteCarrera{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
