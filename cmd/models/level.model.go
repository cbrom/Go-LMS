package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"
)

// Level defines a model for course levels
type Level struct {
	utils.Base
	Name        string `gorm:"type:varchar(100)"`
	CourseID    string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Description string
	Number      int
	UnlockOn    *time.Time
}
