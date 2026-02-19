package main

import (
	"kpo-mini-dz2/infrastructure/repositories"
	"kpo-mini-dz2/presentation/controllers"
	"net/http"

	_ "kpo-mini-dz2/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title My Chi API
// @version 1.0
// @description This is a sample server for my API using Chi
// @termsOfService https://example.com/terms/

// @contact.name API Support
// @contact.url https://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// 1. Инициализация репозиториев
	animalRepo := repositories.NewAnimalRepository()
	enclosureRepo := repositories.NewInMemoryEnclosureRepository()
	//feedingRepo := repositories.NewInMemoryFeedingScheduleRepository()

	// 2. Инициализация контроллеров
	animalHandler := &controllers.AnimalHandler{Repo: animalRepo}
	zooStatsHandler := &controllers.ZooStatisticsHandler{AnimalRepo: animalRepo, EnclosureRepo: enclosureRepo}
	//feedingHandler := &controllers.FeedingHandler{Repo: feedingRepo}

	// 4. API роуты
	r.Route("/api", func(r chi.Router) {
		// Животные
		r.Route("/animals", func(r chi.Router) {
			r.Get("/", animalHandler.GetAll)
			r.Post("/", animalHandler.Create)
			r.Get("/{id}", animalHandler.GetByID)
			r.Delete("/{id}", animalHandler.Delete)
		})
		// Вольеры
		r.Route("/enclosures", func(r chi.Router) {
			r.Get("/", zooStatsHandler.GetAllEnclosures)
		})
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// 6. Запуск сервера
	port := ":3000"
	println("Swagger docs: http://localhost" + port + "/swagger/index.html")
	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
