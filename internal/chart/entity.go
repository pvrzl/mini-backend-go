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
