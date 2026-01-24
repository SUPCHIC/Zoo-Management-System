package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Animal struct {
	ID           uuid.UUID
	Name         string
	Species      Species
	BirthDate    time.Time
	EnclosureID  uuid.UUID
	HealthStatus HealthStatus
	Gender       Gender
	FavoriteFood Food
}

func NewAnimal(
	name string,
	species Species,
	birthDate time.Time,
	enclosureID uuid.UUID,
	healthStatus HealthStatus,
	gender Gender,
	favoriteFood Food,
) (*Animal, error) {
	if name == "" {
		return nil, errors.New("имя не может быть пустым")
	}
	if birthDate.After(time.Now()) {
		return nil, errors.New("дата рождения не может быть из будущего")
	}

	animal := &Animal{
		ID:           uuid.New(),
		Name:         name,
		Species:      species,
		BirthDate:    birthDate,
		EnclosureID:  enclosureID,
		HealthStatus: healthStatus,
		Gender:       gender,
		FavoriteFood: favoriteFood,
	}
	return animal, nil
}

func (a Animal) Feed() {

}
func (a *Animal) Heal() {
	a.HealthStatus = Healthy
}
func (a Animal) Replace(e *Enclosure) {
	a.EnclosureID = e.ID
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type FoodType string

const (
	Meat      = "meat"
	Grass     = "grass"
	Fish      = "fish"
	Fruit     = "fruit"
	Vegetable = "vegetable"
)

type HealthStatus string

const (
	Healthy = "healthy"
	Sick    = "sick"
)

type AnimalType string

const (
	Predator  = "predator"
	Herbivore = "herbivore"
	Omnivore  = "omnivore"
	Aquatic   = "aquatic"
	Avian     = "avian"
)
