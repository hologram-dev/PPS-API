package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type PlanEstudiosUseCase struct {
}

func (peu *PlanEstudiosUseCase) Create(c context.Context, planEstudios domain.PlanEstudios) error {
	db := bootstrap.DB
	err := db.Create(&planEstudios)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (peu *PlanEstudiosUseCase) Fetch(c context.Context) ([]domain.PlanEstudios, error) {
	db := bootstrap.DB
	planEstudios := []domain.PlanEstudios{}
	err := db.Find(&planEstudios)
	if err.Error != nil {
		return nil, err.Error
	}
	return planEstudios, nil
}

func (peu *PlanEstudiosUseCase) FetchById(c context.Context, id uuid.UUID) (domain.PlanEstudios, error) {
	db := bootstrap.DB
	planestudios := domain.PlanEstudios{}
	err := db.Where("id = ?", id).First(&planestudios)
	if err.Error != nil {
		return domain.PlanEstudios{}, err.Error
	}
	return planestudios, nil
}

func (peu *PlanEstudiosUseCase) Update(c context.Context, updatedPlanEstudios domain.PlanEstudios) error {
	db := bootstrap.DB
	if err := db.Model(&updatedPlanEstudios).
		Omit("deleted_at", "created_at").
		Updates(updatedPlanEstudios).Error; err != nil {
		return err
	}
	return nil
}

func (peu *PlanEstudiosUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.PlanEstudios{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
