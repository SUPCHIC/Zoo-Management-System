package controllers

import (
	"encoding/json"
	"kpo-mini-dz2/domain/model"
	RP "kpo-mini-dz2/domain/repositoriesInterfaces"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ZooStatisticsHandler struct {
	AnimalRepo    RP.IAnimalRepository
	EnclosureRepo RP.IEnclosureRepository
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

// GetAllEnclosures godoc
// @Summary Получить все клетки
// @Tags ZooStat
// @Produce json
// @Success 200 {object} model.Enclosure
// @Router /api/zoostat/ [get]
func (h *ZooStatisticsHandler) GetAllEnclosures(w http.ResponseWriter, r *http.Request) {
	enclosures, err := h.EnclosureRepo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enclosures)
}

// GetAnimalsBySpecies godoc
// @Summary Получить всех животных по виду
// @Tags ZooStat
// @Produce json
// @Success 200 {array} model.Animal
// @Router /api/zoostat/{species} [get]
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

// GetEnclousersByType godoc
// @Summary Get enclousers by type
// @Description Get feeding schedule by ID
// @Tags ZooStat
// @Produce json
// @Success 200 {array} model.Enclosure
// @Router /api/zoostat/{type} [get]
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

// GetEnclousersSpacely godoc
// @Summary Get enclousers with free space
// @Tags ZooStat
// @Produce json
// @Success 200 {array} model.Enclosure
// @Router /api/zoostat/space [get]
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

// GetEnclousersByType godoc
// @Summary Get enclousers by type
// @Description Get feeding schedule by ID
// @Tags ZooStat
// @Produce json
// @Success 200 {integer} int "Animal count"
// @Router /api/zoostat/count [get]
func (h *ZooStatisticsHandler) GetAnimalCount(w http.ResponseWriter, r *http.Request) {
	count := h.AnimalRepo.AnimalCount()

	response := map[string]int{
		"animal_count": count,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetEnclousersByType godoc
// @Summary Get enclousers by type
// @Description Get feeding schedule by ID
// @Tags ZooStat
// @Produce json
// @Success 200 {integer} int "Animal count"
// @Router /api/zoostat/count [get]
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
