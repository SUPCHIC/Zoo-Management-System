package model

type Species struct {
	AnimalType AnimalType `json:"animalType"`
	Name       string     `json:"name"`
}

func NewSpecies(animalType AnimalType, name string) *Species {
	return &Species{
		AnimalType: animalType,
		Name:       name,
	}
}
