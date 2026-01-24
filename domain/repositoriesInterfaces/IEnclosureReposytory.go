package repositoriesinterfaces

import (
	"kpo-mini-dz2/domain/model"

	"github.com/google/uuid"
)

type IEnclosureRepository interface {
	Save(enclosure model.Enclosure) error
	FindByID(id uuid.UUID) (*model.Enclosure, error)
	FindAll() ([]model.Enclosure, error)
	FindByType(animalType model.AnimalType) ([]model.Enclosure, error)
	FindWithAvailableSpace(minSpace int) ([]model.Enclosure, error)
	Update(enclosure model.Enclosure) error
	Delete(id uuid.UUID) error
}
