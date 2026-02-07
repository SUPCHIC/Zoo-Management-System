package controllers

import (
	"encoding/json"
	"kpo-mini-dz2/domain/model"
	"kpo-mini-dz2/infrastructure/repositories"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ZooStatisticsHandler struct {
	AnimalRepo    *repositories.InMemoryAnimalRepository
	EnclosureRepo *repositories.InMemoryEnclosureRepository
}

func NewZooStatisticsHandler(
	animalRepo *repositories.InMemoryAnimalRepository,
	enclosureRepo *repositories.InMemoryEnclosureRepository,
) *ZooStatisticsHandler {
	return &ZooStatisticsHandler{
		AnimalRepo:    animalRepo,
		EnclosureRepo: enclosureRepo,
	}
}

func (h *ZooStatisticsHandler) GetAllAnimals(w http.ResponseWriter, r *http.Request) {
	animals, err := h.AnimalRepo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

func (h *ZooStatisticsHandler) GetAllEnclosures(w http.ResponseWriter, r *http.Request) {
	enclosures, err := h.EnclosureRepo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enclosures)
}

func (h *ZooStatisticsHandler) GetAnimalsBySpecies(w http.ResponseWriter, r *http.Request) {
	speciesName := chi.URLParam(r, "species")

	animals, err := h.AnimalRepo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filteredAnimals := make([]model.Animal, 0)
	for _, animal := range animals {
		if animal.Species.Name == speciesName {
			filteredAnimals = append(filteredAnimals, animal)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredAnimals)
}

func (h *ZooStatisticsHandler) GetEnclosuresByType(w http.ResponseWriter, r *http.Request) {
	typeParam := chi.URLParam(r, "type")
	animalType := model.AnimalType(typeParam)

	enclosures, err := h.EnclosureRepo.FindByType(animalType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enclosures)
}

func (h *ZooStatisticsHandler) GetEnclosuresWithAvailableSpace(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр minSpace из query, если есть
	minSpace := 1
	minSpaceStr := r.URL.Query().Get("minSpace")
	if minSpaceStr != "" {
		if val, err := strconv.Atoi(minSpaceStr); err == nil && val > 0 {
			minSpace = val
		}
	}

	enclosures, err := h.EnclosureRepo.FindWithAvailableSpace(minSpace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enclosures)
}

// Дополнительные методы, которые могут пригодиться

func (h *ZooStatisticsHandler) GetAnimalCount(w http.ResponseWriter, r *http.Request) {
	count := h.AnimalRepo.AnimalCount()

	response := map[string]int{
		"animal_count": count,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ZooStatisticsHandler) GetEnclosureByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	enclosure, err := h.EnclosureRepo.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enclosure)
}

func (h *ZooStatisticsHandler) GetAnimalsInEnclosure(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	enclosureID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	enclosure, err := h.EnclosureRepo.FindByID(enclosureID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Получаем всех животных в вольере
	animalsInEnclosure := make([]model.Animal, 0)
	for _, animalID := range enclosure.AnimalsID {
		animal, err := h.AnimalRepo.FindByID(animalID)
		if err == nil && animal != nil {
			animalsInEnclosure = append(animalsInEnclosure, *animal)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animalsInEnclosure)
}
