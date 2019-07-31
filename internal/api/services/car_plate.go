package services

import "do/internal/api/models"

type CarPlateService interface {
	Get(id string) (*models.CarPlate, error)
	GetAll() ([]*models.CarPlate, error)
	Create(*models.CarPlate) error
	Update(*models.CarPlate) error
	Delete(id string) error
}
