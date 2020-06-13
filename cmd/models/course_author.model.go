package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// CourseAuthor defines a model for course authors
type CourseAuthor struct {
	utils.Base
	UserID   string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	CourseID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}
