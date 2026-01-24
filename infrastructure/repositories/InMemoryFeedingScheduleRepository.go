package repositories

import (
	"kpo-mini-dz2/domain/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

type InMemoryFeedingScheduleRepository struct {
	mu        sync.RWMutex
	schedules map[uuid.UUID][]model.FeedingSchedule
}

func NewInMemoryFeedingScheduleRepository() *InMemoryFeedingScheduleRepository {
	return &InMemoryFeedingScheduleRepository{
		schedules: make(map[uuid.UUID][]model.FeedingSchedule),
	}
}

func (r *InMemoryFeedingScheduleRepository) AddSchedule(schedule model.FeedingSchedule) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.schedules[schedule.AnimalID] = append(r.schedules[schedule.AnimalID], schedule)
	return nil
}

func (r *InMemoryFeedingScheduleRepository) GetSchedulesByAnimalID(animalID uuid.UUID) ([]model.FeedingSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.schedules[animalID], nil
}

func (r *InMemoryFeedingScheduleRepository) GetAllSchedules() (map[uuid.UUID][]model.FeedingSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Создаем копию мапы для безопасного возврата
	result := make(map[uuid.UUID][]model.FeedingSchedule)
	for animalID, schedules := range r.schedules {
		result[animalID] = append([]model.FeedingSchedule{}, schedules...)
	}

	return result, nil
}

func (r *InMemoryFeedingScheduleRepository) RemoveSchedule(animalID uuid.UUID, feedingTime time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	schedules, exists := r.schedules[animalID]
	if !exists {
		return nil
	}
	newSchedules := make([]model.FeedingSchedule, 0, len(schedules))
	for _, schedule := range schedules {
		if !schedule.FeedingTime.Equal(feedingTime) {
			newSchedules = append(newSchedules, schedule)
		}
	}

	r.schedules[animalID] = newSchedules
	return nil
}
func (r *InMemoryFeedingScheduleRepository) ClearSchedules(animalID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.schedules, animalID)
	return nil
}

// ClearAll очищает все расписания
func (r *InMemoryFeedingScheduleRepository) ClearAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.schedules = make(map[uuid.UUID][]model.FeedingSchedule)
	return nil
}
