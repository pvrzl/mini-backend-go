package chart

import (
	"lion/pkg/validator"
)

var (
	stringValidator = validator.NewStringValidator()
)

// Chart entity
type Chart struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// ValidateInsert list of insert validation rule
func (c *Chart) ValidateInsert() error {
	return validator.MergeError(
		validator.ValidateString(
			"name",
			c.Name,
			stringValidator.Required,
		),
	)
}

// Favorite entity
type Favorite struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	ChartId int `json:"chart_id"`
}

// ValidateInsert list of insert validation rule
func (f *Favorite) ValidateInsert(us, cs iForeignService) error {
	err := us.IsIDExist(f.UserId)
	if err != nil {
		return err
	}
	err = cs.IsIDExist(f.ChartId)
	if err != nil {
		return err
	}
	return nil
}
