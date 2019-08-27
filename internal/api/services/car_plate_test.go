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

func (m *StorageMock) Update(carplate *models.CarPlate) bool {
	args := m.Called(carplate)
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

func TestGetFailure(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	id := "111"
	mock.On("Get", id).Return((*models.CarPlate)(nil), false)

	carplate, err := service.Get(id)

	mock.AssertExpectations(t)
	assert.Error(t, err)
	assert.Nil(t, carplate)
}

func TestGetAll(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	id1, id2 := "id1", "id2"
	mock.On("GetAll").Return([]*models.CarPlate{
		&models.CarPlate{ID: id1},
		&models.CarPlate{ID: id2},
	})

	carplates := service.GetAll()

	mock.AssertExpectations(t)
	assert.Equal(t, carplates[0].ID, id1)
	assert.Equal(t, carplates[1].ID, id2)
}

func TestAddSuccess(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	carplate := &models.CarPlate{}
	mock.On("Add", carplate).Return(true)

	err := service.Add(carplate)

	mock.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestAddFailure(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	carplate := &models.CarPlate{}
	mock.On("Add", carplate).Return(false)

	err := service.Add(carplate)

	mock.AssertExpectations(t)
	assert.Error(t, err)
}

func TestUpdateSuccess(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	carplate := &models.CarPlate{ID: "id1"}
	mock.On("Update", carplate).Return(true)

	err := service.Update(carplate)

	mock.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestUpdateFailure(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	carplate := &models.CarPlate{ID: "id1"}
	mock.On("Update", carplate).Return(false)

	err := service.Update(carplate)

	mock.AssertExpectations(t)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	mock := new(StorageMock)
	service := NewCarPlateService(mock)
	carplateId := "id1"
	mock.On("Delete", carplateId).Return(false)

	service.Delete(carplateId)

	mock.AssertExpectations(t)
}
