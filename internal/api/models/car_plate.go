package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type CarPlate struct {
	ID   string `json:"id,omitempty" valid:"Required"`
	Name string `json:"name" valid:"Required"`
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

// type Owner struct {
// 	ID      string `json:"id"`
// 	Name    string `json:"name"`
// 	Surname string `json:"surname"`
// }

// type Car struct {
// 	ModelName string `json:"modelName"`
// 	Make      int16  `json:"make"`
// 	ImageUrl  string `json:"imageUrl"`
// }
