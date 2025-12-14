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
	postgres, err := database.OpenPostgres("postgres://app:app@localhost:5432/app?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUserRepository(postgres)
	useCase := usecase.NewCreateUserUsecase(userRepository, 3*time.Second)
	handler := controller.NewHandler(useCase)
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
