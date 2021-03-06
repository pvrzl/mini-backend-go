package main

import (
	"net/http"

	"lion/internal/chart"
	"lion/internal/users"
	"lion/pkg/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	PORT = utils.GetEnv("PORT", ":8000")
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	userRepo := users.NewRepo()

	// router
	r.Mount("/users", users.NewService(users.ServiceConfig{
		Repo:     userRepo,
		Response: new(utils.HTTPJSONResponse),
	}))

	r.Mount("/charts", chart.NewService(chart.ServiceConfig{
		Repo:        chart.NewRepo(),
		UserService: userRepo,
		Response:    new(utils.HTTPJSONResponse),
	}))

	println("serve at", PORT)
	http.ListenAndServe(PORT, r)
}
