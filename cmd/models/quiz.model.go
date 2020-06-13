package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// Quiz defines a model for target quizes
type Quiz struct {
	utils.Base
	Title    string
	TargetID string `sql:"type:uuid" validate:"required,omitempty,uuid"`
}
