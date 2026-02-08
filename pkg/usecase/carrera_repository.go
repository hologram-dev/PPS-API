package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type CarreraUseCase struct{}

func (cc *CarreraUseCase) Create(c context.Context, carrera domain.Carrera) error {
	db := bootstrap.DB
	err := db.Create(&carrera)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (cc *CarreraUseCase) Fetch(c context.Context) ([]domain.Carrera, error) {
	db := bootstrap.DB
	carrera := []domain.Carrera{}
	err := db.Find(&carrera)
	if err.Error != nil {
		return nil, err.Error
	}
	return carrera, nil
}

func (cc *CarreraUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Carrera, error) {
	db := bootstrap.DB
	carrera := domain.Carrera{}
	err := db.Where("id = ?", id).First(&carrera)
	if err.Error != nil {
		return domain.Carrera{}, err.Error
	}
	return carrera, nil
}

func (cc *CarreraUseCase) Update(c context.Context, updatedCarrera domain.Carrera) error {
	db := bootstrap.DB
	if err := db.Model(&updatedCarrera).
		Omit("deleted_at", "created_at").
		Updates(updatedCarrera).Error; err != nil {
		return err
	}
	return nil
}

func (cc *CarreraUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Carrera{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
