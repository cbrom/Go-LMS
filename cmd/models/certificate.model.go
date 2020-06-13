package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// Certificate defines a model for student certificates in a model
type Certificate struct {
	utils.Base
	CourseID      string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	IssuerID      string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	QRCorner      string
	QRScale       int
	Margin        int
	NameOffsetTop int
	FontSize      int
	Message       string
	Active        bool
}
