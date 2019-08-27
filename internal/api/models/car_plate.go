package models

import (
	"fmt"
	"regexp"

	"github.com/astaxie/beego/validation"
)

type AddCarplateResponse struct {
	ID string `json:"id"`
}

type CarPlate struct {
	ID        string `json:"id,omitempty"`
	PlateID   string `json:"plateId"`
	ModelName string `json:"modelName"`
	ModelYear int16  `json:"modelYear,string"`
	Owner     string `json:"owner"`
}

type validationSummary struct {
	Errors []string `json:"errors"`
}

func toValidationSummary(errors []*validation.Error) *validationSummary {
	var errList = []string{}
	for _, e := range errors {
		errMsg := fmt.Sprintf("%s: %s", e.Key, e.Message)
		errList = append(errList, errMsg)
	}

	return &validationSummary{
		Errors: errList,
	}
}

func (c *CarPlate) Validate() (bool, *validationSummary) {
	v := validation.Validation{}

	r, _ := regexp.Compile("^[A-Z]{3}-[0-9]{3}$")
	v.Required(c.ModelName, "modelName").Message("Model name is required")
	v.Required(c.Owner, "owner").Message("Owner is required")
	v.Min(c.ModelYear, 1886, "modelYear").Message("Invalid model year. Should be greater than 1886")
	v.Match(c.PlateID, r, "plateId").Message("Invalid plate format. Should be AAA-000")

	if v.HasErrors() {
		return false, toValidationSummary(v.Errors)
	}
	return true, nil
}
