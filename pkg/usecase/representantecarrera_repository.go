package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type RepresentanteCarreraUseCase struct{}

func (rcu *RepresentanteCarreraUseCase) Create(c context.Context, representanteCarrera domain.RepresentanteCarrera) error {
	db := bootstrap.DB
	err := db.Create(&representanteCarrera)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (rcu *RepresentanteCarreraUseCase) Fetch(c context.Context) ([]domain.RepresentanteCarrera, error) {
	db := bootstrap.DB
	representanteCarreras := []domain.RepresentanteCarrera{}
	err := db.Find(&representanteCarreras)
	if err.Error != nil {
		return nil, err.Error
	}
	return representanteCarreras, nil
}

func (rcu *RepresentanteCarreraUseCase) FetchById(c context.Context, id uuid.UUID) (domain.RepresentanteCarrera, error) {
	db := bootstrap.DB
	representanteCarrera := domain.RepresentanteCarrera{}
	err := db.Where("id = ?", id).First(&representanteCarrera)
	if err.Error != nil {
		return domain.RepresentanteCarrera{}, err.Error
	}
	return representanteCarrera, nil
}

func (rcu *RepresentanteCarreraUseCase) Update(c context.Context, updatedRepresentanteCarrera domain.RepresentanteCarrera) error {
	db := bootstrap.DB
	if err := db.Model(&updatedRepresentanteCarrera).
		Omit("deleted_at", "created_at").
		Updates(updatedRepresentanteCarrera).Error; err != nil {
		return err
	}
	return nil
}

func (rcu *RepresentanteCarreraUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.RepresentanteCarrera{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
