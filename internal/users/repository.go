package users

import (
	"errors"
	"sync"
)

// dummy db/datastore
type repo struct {
	mux   sync.RWMutex
	users []*User
}

var (
	// ErrNotfound occured if no record found in db
	ErrNotfound error = errors.New("records not found")
)

// NewRepo create a new repo
func NewRepo() *repo {
	return new(repo)
}

func (r *repo) Save(user *User) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	length := len(r.users)
	if length == 0 {
		user.Id = 1
	} else {
		user.Id = r.users[length-1].Id + 1
	}

	r.users = append(r.users, user)

	return nil
}

func (r *repo) GetAll() ([]User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	users := make([]User, len(r.users))
	for i := range users {
		users[i] = *r.users[i]
		users[i].Password = ""
	}
	return users, nil
}

func (r *repo) GetByEmail(email string) (*User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	for _, v := range r.users {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, ErrNotfound
}

func (r *repo) Delete(id int) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	idx := r.getIndex(id)
	if idx == -1 {
		return ErrNotfound
	}

	r.removeIndex(idx)
	return nil
}

func (r *repo) removeIndex(index int) {
	r.users = append(r.users[:index], r.users[index+1:]...)
}

func (r *repo) getIndex(id int) int {
	for i, v := range r.users {
		if v.Id == id {
			return i
		}
	}

	return -1
}
