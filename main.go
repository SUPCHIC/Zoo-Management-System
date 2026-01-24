package main

import (
	_ "kpo-mini-dz2/docs"
	"kpo-mini-dz2/infrastructure/repositories"
	"kpo-mini-dz2/presentation/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Animal Service API
// @version 1.0
// @description Это сервер для управления животными в зоопарке.
// @host localhost:3000
// @BasePath /api

// @contact.name API Support
// @contact.url http://example.com/support

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// 1. Инициализация репозиториев
	animalRepo := repositories.NewAnimalRepository()
	enclosureRepo := repositories.NewInMemoryEnclosureRepository()
	feedingRepo := repositories.NewInMemoryFeedingScheduleRepository()

	// 2. Инициализация контроллеров
	animalHandler := &controllers.AnimalHandler{Repo: animalRepo}
	zooStatsHandler := controllers.NewZooStatisticsHandler(animalRepo, enclosureRepo)
	feedingHandler := controllers.NewFeedingHandler(feedingRepo)

	// 3. Роуты для Swagger документации
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	// 4. API роуты
	r.Route("/api", func(r chi.Router) {
		// Животные
		r.Route("/animals", func(r chi.Router) {
			// @Summary Добавить новое животное
			// @Description Создает новое животное в системе
			// @Tags animals
			// @Accept json
			// @Produce json
			// @Param animal body model.Animal true "Данные животного"
			// @Success 201 {object} model.Animal
			// @Failure 400 {object} map[string]string
			// @Router /animals [post]
			r.Post("/", animalHandler.Create)

			// @Summary Получить всех животных
			// @Description Возвращает список всех животных в зоопарке
			// @Tags animals
			// @Produce json
			// @Success 200 {array} model.Animal
			// @Router /animals [get]
			r.Get("/", animalHandler.GetAll)

			// @Summary Получить животное по ID
			// @Description Возвращает информацию о животном по его идентификатору
			// @Tags animals
			// @Produce json
			// @Param id path string true "ID животного"
			// @Success 200 {object} model.Animal
			// @Failure 404 {object} map[string]string
			// @Router /animals/{id} [get]
			r.Get("/{id}", animalHandler.GetByID)

			// @Summary Удалить животное
			// @Description Удаляет животное из системы по его ID
			// @Tags animals
			// @Param id path string true "ID животного"
			// @Success 204
			// @Failure 404 {object} map[string]string
			// @Router /animals/{id} [delete]
			r.Delete("/{id}", animalHandler.Delete)
		})

		// Вольеры
		r.Route("/enclosures", func(r chi.Router) {
			// @Summary Получить все вольеры
			// @Description Возвращает список всех вольеров в зоопарке
			// @Tags enclosures
			// @Produce json
			// @Success 200 {array} model.Enclosure
			// @Router /enclosures [get]
			r.Get("/", zooStatsHandler.GetAllEnclosures)

			// @Summary Получить вольеры с доступным местом
			// @Description Возвращает вольеры, где есть свободное место для животных
			// @Tags enclosures
			// @Produce json
			// @Param minSpace query int false "Минимальное доступное место (по умолчанию 1)"
			// @Success 200 {array} model.Enclosure
			// @Router /enclosures/available [get]
			r.Get("/available", zooStatsHandler.GetEnclosuresWithAvailableSpace)

			// @Summary Получить вольеры по типу
			// @Description Возвращает вольеры определенного типа (хищники, травоядные и т.д.)
			// @Tags enclosures
			// @Produce json
			// @Param type path string true "Тип животных (predator, herbivore, omnivore, aquatic, avian)"
			// @Success 200 {array} model.Enclosure
			// @Router /enclosures/type/{type} [get]
			r.Get("/type/{type}", zooStatsHandler.GetEnclosuresByType)
		})

		// Статистика зоопарка
		r.Route("/statistics", func(r chi.Router) {
			// @Summary Получить всех животных
			// @Description Возвращает статистику по всем животным
			// @Tags statistics
			// @Produce json
			// @Success 200 {array} model.Animal
			// @Router /statistics/animals [get]
			r.Get("/animals", zooStatsHandler.GetAllAnimals)

			// @Summary Получить животных по виду
			// @Description Возвращает животных определенного вида
			// @Tags statistics
			// @Produce json
			// @Param species path string true "Название вида"
			// @Success 200 {array} model.Animal
			// @Router /statistics/animals/species/{species} [get]
			r.Get("/animals/species/{species}", zooStatsHandler.GetAnimalsBySpecies)
		})

		// Кормление
		r.Route("/feeding", func(r chi.Router) {
			// @Summary Добавить расписание кормления
			// @Description Создает новое расписание кормления для животного
			// @Tags feeding
			// @Accept json
			// @Produce json
			// @Param schedule body controllers.AddScheduleRequest true "Данные расписания"
			// @Success 201 {object} map[string]string
			// @Failure 400 {object} map[string]string
			// @Router /feeding/schedule [post]
			r.Post("/schedule", feedingHandler.AddSchedule)

			// @Summary Удалить расписание кормления
			// @Description Удаляет расписание кормления для животного
			// @Tags feeding
			// @Accept json
			// @Produce json
			// @Param schedule body controllers.RemoveScheduleRequest true "Данные для удаления расписания"
			// @Success 200 {object} map[string]string
			// @Failure 400 {object} map[string]string
			// @Router /feeding/schedule [delete]
			r.Delete("/schedule", feedingHandler.RemoveSchedule)

			// @Summary Получить расписание кормления животного
			// @Description Возвращает все расписания кормления для конкретного животного
			// @Tags feeding
			// @Produce json
			// @Param animalID path string true "ID животного"
			// @Success 200 {array} model.FeedingSchedule
			// @Router /feeding/schedule/{animalID} [get]
			r.Get("/schedule/{animalID}", feedingHandler.GetAnimalSchedules)
		})
	})

	// 5. Health check
	// @Summary Проверка здоровья сервиса
	// @Description Проверяет, работает ли сервис
	// @Tags health
	// @Produce plain
	// @Success 200 {string} string "OK"
	// @Router /health [get]
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 6. Запуск сервера
	port := ":3000"
	println("Server started on http://localhost" + port)
	println("Swagger docs: http://localhost" + port + "/swagger/index.html")
	println("API base path: http://localhost" + port + "/api")
	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
