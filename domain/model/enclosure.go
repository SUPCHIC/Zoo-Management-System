package model

import (
	"errors"

	"github.com/google/uuid"
)

type Enclosure struct {
	AnimalsID    []uuid.UUID `json:"animalsID"`
	ID           uuid.UUID   `json:"ID"`
	Type         AnimalType  `json:"type"`
	Size         Size        `json:"size"`
	CurrentCount int         `json:"currentCount"`
	MaxCapacity  int         `json:"maxCapacity"`
}

func NewEnclosure(
	enclosureType AnimalType,
	size Size,
	maxCapacity int,
) (*Enclosure, error) {

	if maxCapacity <= 0 {
		return nil, errors.New("вместимость должна быть больше нуля")
	}

	enclosure := &Enclosure{
		ID:           uuid.New(),
		Type:         enclosureType,
		Size:         size,
		CurrentCount: 0,
		MaxCapacity:  maxCapacity,
		AnimalsID:    []uuid.UUID{},
	}

	return enclosure, nil
}

func (e *Enclosure) AddAnimal(a Animal) {
	e.AnimalsID = append(e.AnimalsID, a.ID)
}
func (e *Enclosure) DeleteAnimal(a Animal) {
	for i, id := range e.AnimalsID {
		if id == a.ID {
			e.AnimalsID = append(e.AnimalsID[:i], e.AnimalsID[i+1:]...)
		}
	}
}

func (en1 Enclosure) ReplaceAnimal(en2 Enclosure, a Animal) {
	en1.DeleteAnimal(a)
	en2.AddAnimal(a)
}
