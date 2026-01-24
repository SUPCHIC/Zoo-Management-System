package repositories

import (
	"errors"
	"kpo-mini-dz2/domain/model"
	"sync"

	"github.com/google/uuid"
)

type InMemoryAnimalRepository struct {
	mu      sync.RWMutex
	animals map[uuid.UUID]model.Animal
}

func NewAnimalRepository() *InMemoryAnimalRepository {
	return &InMemoryAnimalRepository{
		animals: make(map[uuid.UUID]model.Animal),
	}
}

func (r *InMemoryAnimalRepository) Save(animal model.Animal) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.animals[animal.ID] = animal
	return nil
}

func (r *InMemoryAnimalRepository) FindByID(id uuid.UUID) (*model.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	animal, exists := r.animals[id]
	if !exists {
		return nil, errors.New("животное не найдено")
	}

	return &animal, nil
}

func (r *InMemoryAnimalRepository) FindAll() ([]model.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	animals := make([]model.Animal, 0, len(r.animals))
	for _, animal := range r.animals {
		animals = append(animals, animal)
	}
	return animals, nil
}

func (r *InMemoryAnimalRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.animals, id)
	return nil
}

func (r *InMemoryAnimalRepository) AnimalCount() int {
	count := len(r.animals)
	return count
}
