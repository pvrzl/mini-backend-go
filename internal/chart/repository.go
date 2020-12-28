package chart

import (
	"errors"
	"sync"
)

// dummy db/datastore
type repo struct {
	mux    sync.RWMutex
	charts []*Chart
}

var (
	// ErrNotfound occured if no record found in db
	ErrNotfound error = errors.New("records not found")
)

// NewRepo create a new repo
func NewRepo() *repo {
	return new(repo)
}

func (r *repo) Save(chart *Chart) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	length := len(r.charts)
	if length == 0 {
		chart.Id = 1
	} else {
		chart.Id = r.charts[length-1].Id + 1
	}

	r.charts = append(r.charts, chart)

	return nil
}

func (r *repo) getAll() ([]Chart, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	charts := make([]Chart, len(r.charts))
	for i := range charts {
		charts[i] = *r.charts[i]
	}
	return charts, nil
}

func (r *repo) Get(skip int, limit int) ([]Chart, error) {
	charts, err := r.getAll()
	if len(charts) == 0 {
		return charts, err
	}
	if limit > len(charts) {
		limit = len(charts)
	}

	if skip > len(charts) {
		skip = len(charts)
	}

	if skip > len(charts) {
		skip = len(charts)
	}

	total := skip + limit
	if total > len(charts) {
		total = len(charts)
	}

	charts = charts[skip:total]
	return charts, err
}
