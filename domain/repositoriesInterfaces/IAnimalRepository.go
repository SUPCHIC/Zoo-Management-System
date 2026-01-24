package repositoriesinterfaces

import (
	"kpo-mini-dz2/domain/model"

	"github.com/google/uuid"
)

type IAnimalRepository interface {
	Save(animal model.Animal) error
	FindByID(id uuid.UUID) (*model.Animal, error)
	FindAll() ([]model.Animal, error)
	Delete(id uuid.UUID) error
	AnimalCount() int
}
