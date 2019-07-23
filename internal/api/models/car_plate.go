package models

type CarPlate struct {
	ID string `json:"id"`
	// Owner Owner  `json:"owner"`
	// Car   Car    `json:"car"`
}

// type Owner struct {
// 	ID      string `json:"id"`
// 	Name    string `json:"name"`
// 	Surname string `json:"surname"`
// }

// type Car struct {
// 	ModelName string `json:"modelName"`
// 	Make      int16  `json:"make"`
// }

var carPlates []*CarPlate

func addPlate() {
	carPlates = append(carPlates, &CarPlate{
		ID: "GTR 000",
	})
}

func GetCarPlates() []*CarPlate {
	// s := append(s, &CarPlate{
	// 	ID: "GTR 000",
	// })
	addPlate()
	return carPlates
	// return []*CarPlate{
	// 	&CarPlate{
	// 		ID: "GTR 000",
	// 	},
	// 	&CarPlate{
	// 		ID: "HBOs 200",
	// 	}
	// }
}
