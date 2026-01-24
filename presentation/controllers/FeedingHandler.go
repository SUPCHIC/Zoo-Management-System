package controllers

import (
	"encoding/json"
	"kpo-mini-dz2/domain/model"
	"kpo-mini-dz2/infrastructure/repositories"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FeedingHandler struct {
	Repo *repositories.InMemoryFeedingScheduleRepository
}

func NewFeedingHandler(repo *repositories.InMemoryFeedingScheduleRepository) *FeedingHandler {
	return &FeedingHandler{
		Repo: repo,
	}
}

type AddScheduleRequest struct {
	AnimalID    uuid.UUID      `json:"animalId"`
	FeedingTime time.Time      `json:"feedingTime"`
	FoodType    model.FoodType `json:"foodType"`
}

type RemoveScheduleRequest struct {
	AnimalID    uuid.UUID `json:"animalId"`
	FeedingTime time.Time `json:"feedingTime"`
}

func (h *FeedingHandler) AddSchedule(w http.ResponseWriter, r *http.Request) {
	var req AddScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	schedule := model.FeedingSchedule{
		AnimalID:    req.AnimalID,
		FeedingTime: req.FeedingTime,
		FoodType:    req.FoodType,
	}

	err := h.Repo.AddSchedule(schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Schedule added successfully",
	})
}

func (h *FeedingHandler) RemoveSchedule(w http.ResponseWriter, r *http.Request) {
	var req RemoveScheduleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.Repo.RemoveSchedule(req.AnimalID, req.FeedingTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Schedule removed successfully",
	})
}

func (h *FeedingHandler) GetAnimalSchedules(w http.ResponseWriter, r *http.Request) {
	animalIDStr := chi.URLParam(r, "animalID")
	animalID, err := uuid.Parse(animalIDStr)
	if err != nil {
		http.Error(w, "Invalid animal ID", http.StatusBadRequest)
		return
	}

	schedules, err := h.Repo.GetSchedulesByAnimalID(animalID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if schedules == nil {
		schedules = []model.FeedingSchedule{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}
