package storages

import (
	"do/internal/api/models"
	"sync"

	cmap "github.com/orcaman/concurrent-map"
)

type CarPlateStorage interface {
	Get(id string) (*models.CarPlate, bool)
	GetAll() []*models.CarPlate
	Add(m *models.CarPlate) bool
	Update(id string, m *models.CarPlate) bool
	Delete(id string)
}

var once sync.Once
var instance *CarPlateStorageImpl

func NewCarPlateStorage() CarPlateStorage {
	once.Do(func() {
		instance = &CarPlateStorageImpl{
			carplates: cmap.New(),
		}
	})
	return instance
}

type CarPlateStorageImpl struct {
	carplates cmap.ConcurrentMap
}

func (c *CarPlateStorageImpl) Get(id string) (*models.CarPlate, bool) {
	if carplate, found := c.carplates.Get(id); found {
		return carplate.(*models.CarPlate), found
	}
	return nil, false
}

func (c *CarPlateStorageImpl) GetAll() []*models.CarPlate {
	carplatesMap := c.carplates.Items()
	results := make([]*models.CarPlate, 0, len(carplatesMap))

	for _, v := range carplatesMap {
		results = append(results, v.(*models.CarPlate))
	}
	return results
}

func (c *CarPlateStorageImpl) Add(m *models.CarPlate) bool {
	if m.ID == "" {
		return false
	}
	return c.carplates.SetIfAbsent(m.ID, m)
}

func (c *CarPlateStorageImpl) Update(id string, m *models.CarPlate) bool {
	if c.carplates.Has(id) {
		c.carplates.Set(id, m)
		return true
	}

	return false
}

func (c *CarPlateStorageImpl) Delete(id string) {
	c.carplates.Remove(id)
}
