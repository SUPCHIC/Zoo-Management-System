package model

type Species struct {
	AnimalType AnimalType
	Name       string
}

func NewSpecies(animalType AnimalType, name string) *Species {
	return &Species{
		AnimalType: animalType,
		Name:       name,
	}
}
