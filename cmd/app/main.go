package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sinakeshmiri/imcore/api/controller"
	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/infrastructure/database"
	"github.com/sinakeshmiri/imcore/repository"
	"github.com/sinakeshmiri/imcore/usecase"
)

func main() {
	postgres, err := database.OpenPostgres("postgres://app:app@localhost:5435/app?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUserRepository(postgres)
	roleRepository := repository.NewRoleRepository(postgres)
	applicationRepository := repository.NewApplicationRepository(postgres)
	userUseCase := usecase.NewUserUsecase(userRepository, 3*time.Second)
	roleUsecase := usecase.NewRoleUsecase(roleRepository, 3*time.Second)
	applicationUsecase := usecase.NewApplicationUsecase(applicationRepository, 3*time.Second)
	handler := controller.NewHandler(userUseCase, roleUsecase, applicationUsecase)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	strict := api.NewStrictHandlerWithOptions(handler, nil, api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	})
	api.HandlerFromMux(strict, r)
	log.Println("server running on :8080")
	log.Println(http.ListenAndServe(":8080", r))
}
