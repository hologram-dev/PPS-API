package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type ProyectoUseCase struct{}

func (pu *ProyectoUseCase) Create(c context.Context, proyecto domain.Proyecto) error {
	db := bootstrap.DB
	err := db.Create(&proyecto)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (pu *ProyectoUseCase) Fetch(c context.Context) ([]domain.Proyecto, error) {
	db := bootstrap.DB
	proyectos := []domain.Proyecto{}
	err := db.Find(&proyectos)
	if err.Error != nil {
		return nil, err.Error
	}
	return proyectos, nil
}

func (pu *ProyectoUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Proyecto, error) {
	db := bootstrap.DB
	proyecto := domain.Proyecto{}
	err := db.Where("id = ?", id).First(&proyecto)
	if err.Error != nil {
		return domain.Proyecto{}, err.Error
	}
	return proyecto, nil
}

func (pu *ProyectoUseCase) Update(c context.Context, updatedProyecto domain.Proyecto) error {
	db := bootstrap.DB
	if err := db.Model(&updatedProyecto).
		Omit("deleted_at", "created_at").
		Updates(updatedProyecto).Error; err != nil {
		return err
	}
	return nil
}

func (pu *ProyectoUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Proyecto{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
