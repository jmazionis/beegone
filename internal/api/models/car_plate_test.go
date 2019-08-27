package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSuccess(t *testing.T) {
	carplate := &CarPlate{
		PlateID:   "AAA-000",
		ModelName: "Supercar",
		ModelYear: 2010,
		Owner:     "Me",
	}

	ok, validationInfo := carplate.Validate()

	assert.True(t, ok)
	assert.Nil(t, validationInfo)
}

func TestValidateFailure(t *testing.T) {
	carplate := &CarPlate{
		PlateID:   "AAA-00",
		ModelYear: 1000,
	}

	ok, validationInfo := carplate.Validate()

	assert.False(t, ok)
	assert.Equal(t, len(validationInfo.Errors), 4)
}
