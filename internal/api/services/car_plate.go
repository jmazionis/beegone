package services

import (
	"do/internal/api/models"
	"do/internal/api/storages"
	"fmt"
)

type CarPlateService interface {
	Get(id string) (*models.CarPlate, error)
	GetAll() []*models.CarPlate
	Add(*models.CarPlate) error
	Update(*models.CarPlate) error
	Delete(id string)
}

func NewCarPlateService() CarPlateService {
	return &CarPlateServiceImpl{
		storage: storages.NewCarPlateStorage(),
	}
}

type CarPlateServiceImpl struct {
	storage storages.CarPlateStorage
}

func (c *CarPlateServiceImpl) Get(id string) (*models.CarPlate, error) {
	if carplate, ok := c.storage.Get(id); ok {
		return carplate, nil
	}
	return nil, fmt.Errorf("carplate %v does not exist", id)
}

func (c *CarPlateServiceImpl) GetAll() []*models.CarPlate {
	return c.storage.GetAll()
}

func (c *CarPlateServiceImpl) Add(carPlate *models.CarPlate) error {
	if ok := c.storage.Add(carPlate); !ok {
		return fmt.Errorf("carplate %v already exists", carPlate.ID)
	}
	return nil
}

func (c *CarPlateServiceImpl) Update(carPlate *models.CarPlate) error {
	if ok := c.storage.Update(carPlate.ID, carPlate); !ok {
		return fmt.Errorf("carplate %v not found", carPlate.ID)
	}
	return nil
}

func (c *CarPlateServiceImpl) Delete(id string) {
	c.storage.Delete(id)
}
