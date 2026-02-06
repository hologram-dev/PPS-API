package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type ContratoUseCase struct{}

func (cu *ContratoUseCase) Create(c context.Context, contrato domain.Contrato) error {
	db := bootstrap.DB
	err := db.Create(&contrato)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (cu *ContratoUseCase) Fetch(c context.Context) ([]domain.Contrato, error) {
	db := bootstrap.DB
	contratos := []domain.Contrato{}
	err := db.Find(&contratos)
	if err.Error != nil {
		return nil, err.Error
	}
	return contratos, nil
}

func (cu *ContratoUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Contrato, error) {
	db := bootstrap.DB
	contrato := domain.Contrato{}
	err := db.Where("id = ?", id).First(&contrato)
	if err.Error != nil {
		return domain.Contrato{}, err.Error
	}
	return contrato, nil
}

func (cu *ContratoUseCase) Update(c context.Context, updatedContrato domain.Contrato) error {
	db := bootstrap.DB
	if err := db.Model(&updatedContrato).
		Omit("deleted_at", "created_at").
		Updates(updatedContrato).Error; err != nil {
		return err
	}
	return nil
}

func (cu *ContratoUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Contrato{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
