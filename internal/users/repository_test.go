package users

import (
	"testing"
)

func TestCreateRepo(t *testing.T) {
	db := NewRepo()
	err := db.Save(new(User))
	if err != nil {
		t.Error("should not return error on creating and save into repo")
	}

	if len(db.users) != 1 {
		t.Errorf("total data should only 1 since only added once, instead got %d", len(db.users))
	}
}

func TestGetRepo(t *testing.T) {
	db := new(repo)
	err := db.Save(new(User))
	if err != nil {
		t.Error("should not return error on creating repo")
	}

	data, err := db.GetAll()
	if len(data) != 1 {
		t.Errorf("total data should only 1 since only added once, instead got %d", len(db.users))
	}
}

func TestGetByEmail(t *testing.T) {
	db := new(repo)
	err := db.Save(&User{
		Email: "test@test.com",
		Name:  "blabla",
	})
	if err != nil {
		t.Error("should not return error on creating repo")
	}

	// test with invalid value
	_, err = db.GetByEmail("test")
	if err != ErrNotfound {
		t.Error("getbyemail should return err not found because the data is not exist")
	}

	data, err := db.GetByEmail("test@test.com")
	if err != nil {
		t.Error("should not get error because the data is exist")
	}

	if data.Name != "blabla" {
		t.Error("returned data should have same value with the one in repo")
	}
}

func TestDelete(t *testing.T) {
	db := new(repo)
	db.Save(new(User))
	db.Save(new(User))

	if len(db.users) != 2 {
		t.Errorf("total data should only 2 since added twice, instead got %d", len(db.users))
	}

	// test invalid id
	err := db.Delete(3)
	if err != ErrNotfound {
		t.Error("should retur record not found error")
	}

	// test delete valid id
	err = db.Delete(1)
	if err != nil {
		t.Error("should not return any error")
	}

	if db.users[0].Id != 2 {
		t.Error("the leftover data should have id 2")
	}

	db.Save(new(User))
	if db.users[1].Id != 3 {
		t.Error("the new save data should have id 3")
	}

}
