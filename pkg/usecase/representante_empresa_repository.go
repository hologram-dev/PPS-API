package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type RepresentanteEmpresaUseCase struct{}

func (reu *RepresentanteEmpresaUseCase) Create(c context.Context, representanteEmpresa domain.RepresentanteEmpresa) error {
	db := bootstrap.DB
	err := db.Create(&representanteEmpresa)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (reu *RepresentanteEmpresaUseCase) Fetch(c context.Context) ([]domain.RepresentanteEmpresa, error) {
	db := bootstrap.DB
	representanteEmpresa := []domain.RepresentanteEmpresa{}
	err := db.Find(&representanteEmpresa)
	if err.Error != nil {
		return nil, err.Error
	}
	return representanteEmpresa, nil
}

func (reu *RepresentanteEmpresaUseCase) FetchById(c context.Context, id uuid.UUID) (domain.RepresentanteEmpresa, error) {
	db := bootstrap.DB
	representanteEmpresa := domain.RepresentanteEmpresa{}
	err := db.Where("id = ?", id).First(&representanteEmpresa)
	if err.Error != nil {
		return domain.RepresentanteEmpresa{}, err.Error
	}
	return representanteEmpresa, nil
}

func (reu *RepresentanteEmpresaUseCase) Update(c context.Context, updatedRepresentanteEmpresa domain.RepresentanteEmpresa) error {
	db := bootstrap.DB
	if err := db.Model(&updatedRepresentanteEmpresa).
		Omit("deleted_at", "created_at").
		Updates(updatedRepresentanteEmpresa).Error; err != nil {
		return err
	}
	return nil
}

func (reu *RepresentanteEmpresaUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.RepresentanteEmpresa{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
