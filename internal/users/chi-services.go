package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var (
	// ErrNothingToUpdate occured when nothing to update
	ErrNothingToUpdate error = errors.New("nothing to update")
	// ErrInvalidEmailOrPassword occured when user inserting invalid email or password
	ErrInvalidEmailOrPassword error = errors.New("invalid email or password")
)

type (
	iresponse interface {
		Resp(http.ResponseWriter, int, interface{})
		Err(http.ResponseWriter, int, error)
	}

	irepo interface {
		Save(*User) error
		GetAll() ([]User, error)
		Delete(id int) error
		GetByEmail(string) (*User, error)
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
	r.Post("/", svc.Insert)
	r.Get("/", svc.GetAll)
	r.Delete("/{id}", svc.Delete)
	r.Post("/auth", svc.Auth)
	return r
}

func (s service) Insert(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = user.ValidateInsert()
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = user.EncryptPassword()
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}
	err = s.repo.Save(user)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	s.response.Resp(w, http.StatusCreated, http.StatusText(http.StatusCreated))

}

func (s service) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := s.repo.GetAll()
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}
	s.response.Resp(w, http.StatusOK, users)
}

func (s service) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = s.repo.Delete(int(id))
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	s.response.Resp(w, http.StatusOK, http.StatusText(http.StatusOK))

}

func (s service) Auth(w http.ResponseWriter, r *http.Request) {
	form := new(User)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(form)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = form.ValidateAuth()
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, err)
		return
	}

	user, err := s.repo.GetByEmail(form.Email)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, ErrInvalidEmailOrPassword)
		return
	}

	err = user.ComparePassword(form.Password)
	if err != nil {
		s.response.Err(w, http.StatusBadRequest, ErrInvalidEmailOrPassword)
		return
	}

	s.response.Resp(w, http.StatusOK, "dummy token")

}
