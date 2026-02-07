package model

type Food struct {
	FoodType FoodType `json:"foodType"`
	Name     string   `json:"name"`
}

func NewFood(foodType FoodType, name string) *Food {
	return &Food{
		FoodType: foodType,
		Name:     name,
	}
}
