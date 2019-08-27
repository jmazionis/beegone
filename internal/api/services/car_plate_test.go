package services

import (
	"testing"

	"github.com/ICanHaz/beegone/internal/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StorageMock struct {
	mock.Mock
}

func (m *StorageMock) Get(id string) (*models.CarPlate, bool) {
	args := m.Called(id)
	return args.Get(0).(*models.CarPlate), args.Bool(1)
}

func (m *StorageMock) GetAll() []*models.CarPlate {
	args := m.Called()
	return args.Get(0).([]*models.CarPlate)
}

func (m *StorageMock) Add(carplate *models.CarPlate) bool {
	args := m.Called(carplate)
	return args.Bool(0)
}

func (m *StorageMock) Update(id string, carplate *models.CarPlate) bool {
	args := m.Called(id, carplate)
	return args.Bool(0)
}

func (m *StorageMock) Delete(id string) {
	m.Called(id)
}

func (m *StorageMock) Reset() {
	m.Called()
}

func TestGetSuccess(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	id := "111"
	mock.On("Get", id).Return(&models.CarPlate{ID: id}, true)

	carplate, err := service.Get(id)

	mock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, carplate.ID, id)
}

// func TestGetFailure(t *testing.T) {
// 	mock := new(StorageMock)
// 	service := NewCarPlateService(mock)
// 	id := "111"
// 	mock.On("Get", id).Return(nil, false)

// 	_, err := service.Get(id)

// 	mock.AssertExpectations(t)
// 	assert.Error(t, err)
// }
