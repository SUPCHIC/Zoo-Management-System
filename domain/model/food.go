package model

type Food struct {
	FoodType FoodType
	Name     string
}

func NewFood(foodType FoodType, name string) *Food {
	return &Food{
		FoodType: foodType,
		Name:     name,
	}
}
