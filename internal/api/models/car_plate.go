package models

import (
	"github.com/astaxie/beego/validation"
)

type CarPlate struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name" valid:"Required"`
	// 	Owner Owner  `json:"owner"`
	// 	Car   Car    `json:"car"`
}

func (c *CarPlate) Validate(v *validation.Validation) error {
	validation := validation.Validation{}
	_, err := validation.Valid(c)

	if err != nil {
		// handle error
		return err
	}

	return nil
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
