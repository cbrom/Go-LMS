package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"
)

// Course is a model for Courses table
type Course struct {
	utils.Base
	Name                string `gorm:"type:varchar(100);unique" `
	EndsAt              *time.Time
	Description         string
	EnableLeadboard     bool
	PublicSignup        bool
	Featured            bool
	About               string `gorm:"type:varchar(100)"`
	ProgressionBehavior string `gorm:"type:varchar(100)"`
	ProgressionLimit    int
}
