package services

import (
	"fmt"

	"github.com/ICanHaz/beegone/internal/api/models"
	"github.com/ICanHaz/beegone/internal/api/storages"
)

type CarPlateServicer interface {
	Get(id string) (*models.CarPlate, error)
	GetAll() []*models.CarPlate
	Add(*models.CarPlate) error
	Update(*models.CarPlate) error
	Delete(id string)
}

func NewCarPlateService(carplateStorage storages.CarPlateStorager) CarPlateServicer {
	return &CarPlateService{
		storage: carplateStorage,
	}
}

type CarPlateService struct {
	storage storages.CarPlateStorager
}

func (c *CarPlateService) Get(id string) (*models.CarPlate, error) {
	if carplate, ok := c.storage.Get(id); ok {
		return carplate, nil
	}
	return nil, fmt.Errorf("carplate %v does not exist", id)
}

func (c *CarPlateService) GetAll() []*models.CarPlate {
	return c.storage.GetAll()
}

func (c *CarPlateService) Add(carPlate *models.CarPlate) error {

	if ok := c.storage.Add(carPlate); !ok {
		return fmt.Errorf("carplate %v already exists", carPlate.ID)
	}
	return nil
}

func (c *CarPlateService) Update(carPlate *models.CarPlate) error {
	if ok := c.storage.Update(carPlate.ID, carPlate); !ok {
		return fmt.Errorf("carplate %v not found", carPlate.ID)
	}
	return nil
}

func (c *CarPlateService) Delete(id string) {
	c.storage.Delete(id)
}
