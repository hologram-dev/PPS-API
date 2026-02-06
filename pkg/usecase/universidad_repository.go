package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type UniversidadUseCase struct{}

func (uu *UniversidadUseCase) Create(c context.Context, universidad domain.Universidad) error {
	db := bootstrap.DB
	err := db.Create(&universidad)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (uu *UniversidadUseCase) Fetch(c context.Context) ([]domain.Universidad, error) {
	db := bootstrap.DB
	universidades := []domain.Universidad{}
	err := db.Find(&universidades)
	if err.Error != nil {
		return nil, err.Error
	}
	return universidades, nil
}

func (uu *UniversidadUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Universidad, error) {
	db := bootstrap.DB
	universidad := domain.Universidad{}
	err := db.Where("id = ?", id).First(&universidad)
	if err.Error != nil {
		return domain.Universidad{}, err.Error
	}
	return universidad, nil
}

func (uu *UniversidadUseCase) Update(c context.Context, updatedUniversidad domain.Universidad) error {
	db := bootstrap.DB
	if err := db.Model(&updatedUniversidad).
		Omit("deleted_at", "created_at").
		Updates(updatedUniversidad).Error; err != nil {
		return err
	}
	return nil
}

func (uu *UniversidadUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Universidad{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
