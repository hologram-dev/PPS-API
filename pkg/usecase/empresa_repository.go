package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type EmpresaUseCase struct{}

func (eu *EmpresaUseCase) Create(c context.Context, empresa domain.Empresa) error {
	db := bootstrap.DB
	err := db.Create(&empresa)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *EmpresaUseCase) Fetch(c context.Context) ([]domain.Empresa, error) {
	db := bootstrap.DB
	entity := []domain.Empresa{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *EmpresaUseCase) FetchById(c context.Context, id int) (domain.Empresa, error) {
	db := bootstrap.DB
	empresa := domain.Empresa{}
	err := db.Where("id = ?", id).First(&empresa)
	if err.Error != nil {
		return domain.Empresa{}, err.Error
	}
	return empresa, nil
}

func (eu *EmpresaUseCase) Update(c context.Context, updatedEmpresa domain.Empresa) error {
	db := bootstrap.DB
	if err := db.Model(&updatedEmpresa).
		Omit("deleted_at", "created_at").
		Updates(updatedEmpresa).Error; err != nil {
		return err
	}
	return nil
}

func (eu *EmpresaUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Empresa{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
