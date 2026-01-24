package controllers

import (
	"encoding/json"
	"kpo-mini-dz2/domain/model"
	RP "kpo-mini-dz2/domain/repositoriesInterfaces"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AnimalHandler struct {
	Repo RP.IAnimalRepository
}

// Create godoc
// @Summary Добавить животное
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body model.Animal true "Animal Data"
// @Success 201 {object} model.Animal
// @Router /animals [post]
func (h *AnimalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newAnimal model.Animal

	err := json.NewDecoder(r.Body).Decode(&newAnimal)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.Repo.Save(newAnimal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAnimal)
}

// GetAll godoc
// @Summary Получить всех животных
// @Tags animals
// @Produce json
// @Success 200 {array} model.Animal
// @Router /animals [get]
func (h *AnimalHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	animals, err := h.Repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

// GetByID godoc
// @Summary Получить животное по ID
// @Tags animals
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {object} model.Animal
// @Router /animals/{id} [get]
func (h *AnimalHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	animal, err := h.Repo.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

// Delete godoc
// @Summary Удалить животное
// @Tags animals
// @Param id path string true "Animal ID"
// @Success 204
// @Router /animals/{id} [delete]
func (h *AnimalHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
