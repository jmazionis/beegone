package services

import "do/internal/api/models"

type CarPlateService interface {
	// Get(id string) (*models.CarPlate, error)
	GetAll() []*models.CarPlate
	// Create(*models.CarPlate) error
	// Update(*models.CarPlate) error
	// Delete(id string) error
}

func NewCarPlateService() CarPlateService {
	return &CarServiceImpl{}
}

type CarServiceImpl struct {
}

func (c *CarServiceImpl) GetAll() []*models.CarPlate {
	return models.GetCarPlates()
}
