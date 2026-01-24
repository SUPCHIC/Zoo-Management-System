package services

import (
	"kpo-mini-dz2/domain/model"
	RP "kpo-mini-dz2/domain/repositoriesInterfaces"
	"time"

	"github.com/google/uuid"
)

type FeedingService struct {
	repo RP.IFeedingScheduleRepository
}

func NewFeedingService(repo RP.IFeedingScheduleRepository) *FeedingService {
	return &FeedingService{repo: repo}
}

func (s *FeedingService) AddFeedingSchedule(animalID uuid.UUID, feedingTime time.Time, foodType model.FoodType) error {
	schedule := model.FeedingSchedule{
		AnimalID:    animalID,
		FeedingTime: feedingTime,
		FoodType:    foodType,
	}

	return s.repo.AddSchedule(schedule)
}

func (s *FeedingService) RemoveFeedingSchedule(animalID uuid.UUID, feedingTime time.Time) error {
	return s.repo.RemoveSchedule(animalID, feedingTime)
}

func (s *FeedingService) GetAnimalSchedules(animalID uuid.UUID) ([]model.FeedingSchedule, error) {
	return s.repo.GetSchedulesByAnimalID(animalID)
}
