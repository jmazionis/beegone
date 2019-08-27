package storages

import (
	"github.com/ICanHaz/beegone/internal/api/models"

	cmap "github.com/orcaman/concurrent-map"
)

type CarPlateStorager interface {
	Get(id string) (*models.CarPlate, bool)
	GetAll() []*models.CarPlate
	Add(m *models.CarPlate) bool
	Update(m *models.CarPlate) bool
	Delete(id string)
	Reset()
}

var carplateStorage *CarPlateStorage

func init() {
	carplateStorage = &CarPlateStorage{
		carplates: cmap.New(),
	}
}

// A "thread" safe map-based storage of type string:Anything.
func CarPlateDb() CarPlateStorager {
	return carplateStorage
}

type CarPlateStorage struct {
	carplates cmap.ConcurrentMap
}

func (c *CarPlateStorage) Get(id string) (*models.CarPlate, bool) {
	if carplate, found := c.carplates.Get(id); found {
		return carplate.(*models.CarPlate), found
	}
	return nil, false
}

func (c *CarPlateStorage) GetAll() []*models.CarPlate {
	carplatesMap := c.carplates.Items()
	results := make([]*models.CarPlate, 0, len(carplatesMap))

	for _, v := range carplatesMap {
		results = append(results, v.(*models.CarPlate))
	}
	return results
}

func (c *CarPlateStorage) Add(m *models.CarPlate) bool {
	if m.ID == "" {
		return false
	}
	return c.carplates.SetIfAbsent(m.ID, m)
}

func (c *CarPlateStorage) Update(m *models.CarPlate) bool {
	if c.carplates.Has(m.ID) {
		c.carplates.Set(m.ID, m)
		return true
	}
	return false
}

func (c *CarPlateStorage) Delete(id string) {
	c.carplates.Remove(id)
}

func (c *CarPlateStorage) Reset() {
	c.carplates = cmap.New()
}
