package services

import (
	"kpo-mini-dz2/domain/model"
	RP "kpo-mini-dz2/domain/repositoriesInterfaces"
)

/*
ZooStatisticsService - функции:
вывод всех животных
вывод всех клеток с животными
вывод по видам
*/
type ZooStatisticsService struct {
	AnimalRepo    RP.IAnimalRepository
	EnclosureRepo RP.IEnclosureRepository
}

func NewZooStatisticsService(animalRepo RP.IAnimalRepository, enclousureRepo RP.IEnclosureRepository) *ZooStatisticsService {
	return &ZooStatisticsService{AnimalRepo: animalRepo, EnclosureRepo: enclousureRepo}
}

func GetAllAnimals(z *ZooStatisticsService) ([]model.Animal, error) {
	return z.AnimalRepo.FindAll()
}

func GetAllEnclousure(z ZooStatisticsService) ([]model.Enclosure, error) { //reedit
	return z.EnclosureRepo.FindAll()
}

func GetAnimalBySpecies(z *ZooStatisticsService, s model.Species) ([]model.Animal, error) {
	animals, _ := z.AnimalRepo.FindAll()
	AnimalForSpecies := make([]model.Animal, 0)
	for _, animal := range animals {
		if animal.Species == s {
			AnimalForSpecies = append(AnimalForSpecies, animal)
		}
	}
	return AnimalForSpecies, nil
}
