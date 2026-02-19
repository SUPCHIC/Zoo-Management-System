package controllers

import (
	"encoding/json"
	"kpo-mini-dz2/domain/model"
	RP "kpo-mini-dz2/domain/repositoriesInterfaces"
	"kpo-mini-dz2/infrastructure/repositories"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FeedingHandler struct {
	Repo RP.IFeedingScheduleRepository
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

// AddSchedule godoc
// @Summary Add a new feeding schedule
// @Description Adds a new feeding schedule for an animal with specified feeding time and food type
// @Tags feeding_schedule
// @Accept json
// @Produce json
// @Param schedule body AddScheduleRequest true "Feeding schedule information"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /api/schedules [post]
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

// DeleteSchedule godoc
// @Summary Delete feeding schedule
// @Description Delete schedule by animalID and time from repository
// @Tags feeding_schedule
// @Accept json
// @Produce json
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /api/schedules [delete]
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

// GetSchedule godoc
// @Summary Get feeding schedule by ID
// @Description Get feeding schedule by ID
// @Tags feeding_schedule
// @Accept json
// @Produce json
// @Success 201 {object} model.FeedingSchedule
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /api/schedules [get]
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
