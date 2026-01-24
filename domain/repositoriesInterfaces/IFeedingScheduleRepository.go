package repositoriesinterfaces

import (
	"kpo-mini-dz2/domain/model"
	"time"

	"github.com/google/uuid"
)

type IFeedingScheduleRepository interface {
	AddSchedule(schedule model.FeedingSchedule) error
	GetSchedulesByAnimalID(animalID uuid.UUID) ([]model.FeedingSchedule, error)
	GetAllSchedules() (map[uuid.UUID][]model.FeedingSchedule, error)
	RemoveSchedule(animalID uuid.UUID, feedingTime time.Time) error
	ClearSchedules(animalID uuid.UUID) error
}
