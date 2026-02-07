package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"

	"github.com/google/uuid"
)

type PostulacionEstadoUseCase struct{}

func (peu *PostulacionEstadoUseCase) Create(c context.Context, postulacionEstado domain.PostulacionEstado) error {
	db := bootstrap.DB
	err := db.Create(&postulacionEstado)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (peu *PostulacionEstadoUseCase) Fetch(c context.Context) ([]domain.PostulacionEstado, error) {
	db := bootstrap.DB
	postulacionEstado := []domain.PostulacionEstado{}
	err := db.Find(&postulacionEstado)
	if err.Error != nil {
		return nil, err.Error
	}
	return postulacionEstado, nil
}

func (peu *PostulacionEstadoUseCase) FetchById(c context.Context, id uuid.UUID) (domain.PostulacionEstado, error) {
	db := bootstrap.DB
	postulacionEstado := domain.PostulacionEstado{}
	err := db.Where("id = ?", id).First(&postulacionEstado)
	if err.Error != nil {
		return domain.PostulacionEstado{}, err.Error
	}
	return postulacionEstado, nil
}

func (peu *PostulacionEstadoUseCase) Update(c context.Context, updatedPostulacionEstado domain.PostulacionEstado) error {
	db := bootstrap.DB
	if err := db.Model(&updatedPostulacionEstado).
		Omit("deleted_at", "created_at").
		Updates(updatedPostulacionEstado).Error; err != nil {
		return err
	}
	return nil
}

func (peu *PostulacionEstadoUseCase) Delete(c context.Context, id uuid.UUID) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.PostulacionEstado{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
