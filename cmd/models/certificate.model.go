package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// Certificate defines a model for student certificates in a model
type Certificate struct {
	utils.Base
	CourseID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	IssuerID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Message  string
	Active   bool
}
