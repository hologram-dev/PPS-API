package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type PostulacionUseCase struct{}

func (pu *PostulacionUseCase) Create(c context.Context, postulacion domain.Postulacion) error {
	db := bootstrap.DB
	err := db.Create(&postulacion)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (pu *PostulacionUseCase) Fetch(c context.Context) ([]domain.Postulacion, error) {
	db := bootstrap.DB
	postulaciones := []domain.Postulacion{}
	err := db.Find(&postulaciones)
	if err.Error != nil {
		return nil, err.Error
	}
	return postulaciones, nil
}

func (pu *PostulacionUseCase) FetchById(c context.Context, id uuid.UUID) (domain.Postulacion, error) {
	db := bootstrap.DB
	postulacion := domain.Postulacion{}
	err := db.Where("id = ?", id).First(&postulacion)
	if err.Error != nil {
		return domain.Postulacion{}, err.Error
	}
	return postulacion, nil
}

func (pu *PostulacionUseCase) Update(c context.Context, updatedPostulacion domain.Postulacion) error {
	db := bootstrap.DB
	if err := db.Model(&updatedPostulacion).
		Omit("deleted_at", "created_at").
		Updates(updatedPostulacion).Error; err != nil {
		return err
	}
	return nil
}

func (pu *PostulacionUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Postulacion{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
