package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	api "github.com/sinakeshmiri/imcore/api/generated"
	httpapi "github.com/sinakeshmiri/imcore/api/handler"

	appshttp "github.com/sinakeshmiri/imcore/internal/applications/http"
	appsrepo "github.com/sinakeshmiri/imcore/internal/applications/repository"
	appsuc "github.com/sinakeshmiri/imcore/internal/applications/usecase"

	roleshttp "github.com/sinakeshmiri/imcore/internal/roles/http"
	rolesrepo "github.com/sinakeshmiri/imcore/internal/roles/repository"
	rolesuc "github.com/sinakeshmiri/imcore/internal/roles/usecase"

	usershttp "github.com/sinakeshmiri/imcore/internal/users/http"
	usersrepo "github.com/sinakeshmiri/imcore/internal/users/repository"
	usersuc "github.com/sinakeshmiri/imcore/internal/users/usecase"

	"github.com/sinakeshmiri/imcore/infrastructure/database"
)

func main() {
	db, err := database.OpenPostgres("postgres://app:app@localhost:5435/app?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	timeout := 3 * time.Second
	userRepo := usersrepo.NewUserRepository(db)
	roleRepo := rolesrepo.NewRoleRepository(db)
	appRepo := appsrepo.NewApplicationRepository(db)
	userUC := usersuc.NewUserUsecase(userRepo, timeout)
	roleUC := rolesuc.NewRoleUsecase(roleRepo, timeout)
	appUC := appsuc.NewApplicationUsecase(appRepo, timeout)

	userHandler := usershttp.NewHandler(userUC)
	roleHandler := roleshttp.NewHandler(roleUC)
	appHandler := appshttp.NewHandler(appUC)

	root := &httpapi.APIHandler{
		Users:        userHandler,
		Roles:        roleHandler,
		Applications: appHandler,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	strict := api.NewStrictHandlerWithOptions(root, nil, api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	})

	api.HandlerFromMux(strict, r)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
