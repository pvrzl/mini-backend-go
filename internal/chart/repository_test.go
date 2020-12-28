package chart

import (
	"testing"
)

func TestCreateRepo(t *testing.T) {
	db := NewRepo()
	err := db.Save(new(Chart))
	if err != nil {
		t.Error("should not return error on creating and save into repo")
	}

	if len(db.charts) != 1 {
		t.Errorf("total data should only 1 since only added once, instead got %d", len(db.charts))
	}
}

func TestGetRepo(t *testing.T) {
	db := new(repo)
	err := db.Save(new(Chart))
	if err != nil {
		t.Error("should not return error on creating repo")
	}

	data, err := db.getAll()
	if len(data) != 1 {
		t.Errorf("total data should only 1 since only added once, instead got %d", len(db.charts))
	}

	data, err = db.Get(0, 1)
	if len(data) != 1 {
		t.Errorf("total data should only 1 since only added once, instead got %d", len(data))
	}

	data, err = db.Get(0, 0)
	if len(data) != 0 {
		t.Errorf("total data should only 0, instead got %d", len(data))
	}
}
