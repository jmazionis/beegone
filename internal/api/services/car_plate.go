package services

import (
	"do/internal/api/models"
	"do/internal/api/storages"
)

type CarPlateService interface {
	// Get(id string) (*models.CarPlate, error)
	GetAll() []*models.CarPlate
	Add(*models.CarPlate) bool
	// Update(*models.CarPlate) error
	// Delete(id string) error
}

func NewCarPlateService() CarPlateService {
	return &CarServiceImpl{
		storage: storages.NewCarPlateStorage(),
	}
}

type CarServiceImpl struct {
	storage storages.CarPlateStorage
}

func (c *CarServiceImpl) GetAll() []*models.CarPlate {
	return c.storage.GetAll()
}

func (c *CarServiceImpl) Add(carPlate *models.CarPlate) bool {
	ok := c.storage.Add(carPlate)
	return ok
}
