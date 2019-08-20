package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type CarPlate struct {
	ID        string `json:"id,omitempty" valid:"Required"`
	PlateID   string `json:"plateId" valid:"Required;Match([A-Z]{3}-[0-9]{3})"`
	ModelName string `json:"modelName" valid:"Required"`
	ModelYear int16  `json:"modelYear" valid:"Required;Min(1886)"`
	Owner     string `json:"owner" valid:"Required"`
}

type validationSummary struct {
	Errors []string `json:"errors"`
}

func toValidationSummary(errors []*validation.Error) *validationSummary {
	var errList = []string{}
	for _, e := range errors {
		errMsg := fmt.Sprintf("%s field: %s", e.Field, e.Message)
		errList = append(errList, errMsg)
	}

	return &validationSummary{
		Errors: errList,
	}
}

func (c *CarPlate) Validate() (bool, *validationSummary) {
	v := validation.Validation{}
	ok, err := v.Valid(c)

	if err != nil {
		beego.Error(err)
		return false, nil
	}

	return ok, toValidationSummary(v.Errors)
}
