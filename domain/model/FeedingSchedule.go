package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type FeedingSchedule struct {
	AnimalID    uuid.UUID
	FeedingTime time.Time
	FoodType    FoodType
}

func NewFeedingSchedule(
	animalID uuid.UUID,
	feedingTime time.Time,
	foodType FoodType,
) (*FeedingSchedule, error) {
	if animalID == uuid.Nil {
		return nil, errors.New("ID не может быть пустым")
	}

	if feedingTime.Before(time.Now()) {
		return nil, errors.New("времч не может быть в прошллом")
	}

	schedule := &FeedingSchedule{
		AnimalID:    animalID,
		FeedingTime: feedingTime,
		FoodType:    foodType,
	}

	return schedule, nil
}

func (f FeedingSchedule) ChangeSchedule(newTime time.Time) error {
	if newTime.Before(time.Now()) {
		return errors.New("время кормления не может быть в прошлом")
	}

	f.FeedingTime = newTime
	return nil
}
func (f FeedingSchedule) PingExecution() {

}
