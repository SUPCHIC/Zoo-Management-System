package model

import (
	"errors"

	"github.com/google/uuid"
)

type Enclosure struct {
	AnimalsIDs   []uuid.UUID
	ID           uuid.UUID
	Type         AnimalType
	Size         Size
	CurrentCount int
	MaxCapacity  int
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
		AnimalsIDs:   []uuid.UUID{},
	}

	return enclosure, nil
}

func (e *Enclosure) AddAnimal(a Animal) {
	e.AnimalsIDs = append(e.AnimalsIDs, a.ID)
}
func (e *Enclosure) DeleteAnimal(a Animal) {
	for i, id := range e.AnimalsIDs {
		if id == a.ID {
			e.AnimalsIDs = append(e.AnimalsIDs[:i], e.AnimalsIDs[i+1:]...)
		}
	}
}

func (en1 Enclosure) ReplaceAnimal(en2 Enclosure, a Animal) {
	en1.DeleteAnimal(a)
	en2.AddAnimal(a)
}
