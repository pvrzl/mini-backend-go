package chart

import (
	"encoding/json"
	"lion/pkg/middleware"
	"lion/pkg/utils"
	"net/http"

	"github.com/go-chi/chi"
)

type (
	iresponse interface {
		Resp(http.ResponseWriter, int, interface{})
		Err(http.ResponseWriter, int, error)
	}

	irepo interface {
		Save(*Chart) error
		Get(skip, limit int) ([]Chart, error)
	}

	service struct {
		repo     irepo
		response iresponse
	}
	// ServiceConfig is a config for service
	ServiceConfig struct {
		Repo     irepo
		Response iresponse
	}

	genericJSON map[string]interface{}
)

// NewService return a new service
func NewService(cfg ServiceConfig) http.Handler {
	svc := service{
		repo:     cfg.Repo,
		response: cfg.Response,
	}

	r := chi.NewRouter()
	r.Use(middleware.MustLogin)
	r.Post("/", svc.Insert)
	r.Get("/", svc.GetAll)
	return r
}

func (s service) Insert(w http.ResponseWriter, r *http.Request) {
	chart := new(Chart)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(chart)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = chart.ValidateInsert()
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = s.repo.Save(chart)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	s.response.Resp(w, http.StatusCreated, http.StatusText(http.StatusCreated))

}

func (s service) GetAll(w http.ResponseWriter, r *http.Request) {
	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")
	users, err := s.repo.Get(utils.StringToIntWithDefault(skip, 0), utils.StringToIntWithDefault(limit, 10))
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}
	s.response.Resp(w, http.StatusOK, users)
}
