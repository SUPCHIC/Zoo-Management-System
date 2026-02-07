package repositories

import (
	"errors"
	"kpo-mini-dz2/domain/model"
	"sync"

	"github.com/google/uuid"
)

type InMemoryEnclosureRepository struct {
	mu         sync.RWMutex
	enclosures map[uuid.UUID]model.Enclosure
}

func NewInMemoryEnclosureRepository() *InMemoryEnclosureRepository {
	return &InMemoryEnclosureRepository{
		enclosures: make(map[uuid.UUID]model.Enclosure),
	}
}

func (r *InMemoryEnclosureRepository) Save(enclosure model.Enclosure) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.enclosures[enclosure.ID] = enclosure
	return nil
}

func (r *InMemoryEnclosureRepository) FindByID(id uuid.UUID) (*model.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	enclosure, exists := r.enclosures[id]
	if !exists {
		return nil, errors.New("вольер не найден")
	}

	// Возвращаем копию
	return &enclosure, nil
}

// FindAll - возвращает все вольеры
func (r *InMemoryEnclosureRepository) FindAll() ([]model.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	enclosures := make([]model.Enclosure, 0, len(r.enclosures))
	for _, enclosure := range r.enclosures {
		enclosures = append(enclosures, enclosure)
	}

	return enclosures, nil
}

// FindByType - ищет вольеры по типу
func (r *InMemoryEnclosureRepository) FindByType(animalType model.AnimalType) ([]model.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []model.Enclosure
	for _, enclosure := range r.enclosures {
		if enclosure.Type == animalType {
			result = append(result, enclosure)
		}
	}

	return result, nil
}

// FindWithAvailableSpace - ищет вольеры с доступным местом
func (r *InMemoryEnclosureRepository) FindWithAvailableSpace(minSpace int) ([]model.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []model.Enclosure
	for _, enclosure := range r.enclosures {
		available := enclosure.MaxCapacity - enclosure.CurrentCount
		if available >= minSpace {
			result = append(result, enclosure)
		}
	}

	return result, nil
}

// Update - обновляет вольер
func (r *InMemoryEnclosureRepository) Update(enclosure model.Enclosure) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.enclosures[enclosure.ID]; !exists {
		return errors.New("вольер не найден")
	}

	r.enclosures[enclosure.ID] = enclosure
	return nil
}

// Delete - удаляет вольер
func (r *InMemoryEnclosureRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.enclosures[id]; !exists {
		return errors.New("вольер не найден")
	}

	delete(r.enclosures, id)
	return nil
}

/*
// SeedWithSampleData - заполняет репозиторий тестовыми данными
func (r *InMemoryEnclosureRepository) SeedWithSampleData() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Создаем тестовые вольеры
	sampleEnclosures := []entities.Enclosure{
		{
			ID:           uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Type:         enums.Predator,
			Size:         enums.Large,
			CurrentCount: 2,
			MaxCapacity:  5,
		},
		{
			ID:           uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Type:         enums.Herbivore,
			Size:         enums.Large,
			CurrentCount: 10,
			MaxCapacity:  15,
		},
		{
			ID:           uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			Type:         enums.Bird,
			Size:         enums.Medium,
			CurrentCount: 5,
			MaxCapacity:  8,
		},
		{
			ID:           uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			Type:         enums.Aquarium,
			Size:         enums.Small,
			CurrentCount: 3,
			MaxCapacity:  5,
		},
		{
			ID:           uuid.MustParse("00000000-0000-0000-0000-000000000005"),
			Type:         model.Quarantine,
			Size:         model.Small,
			CurrentCount: 1,
			MaxCapacity:  3,
		},
	}

	for _, enclosure := range sampleEnclosures {
		r.enclosures[enclosure.ID] = enclosure
	}

	return nil
}
*/

// Clear - очищает репозиторий (для тестов)
func (r *InMemoryEnclosureRepository) Clear() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.enclosures = make(map[uuid.UUID]model.Enclosure)
	return nil
}
