package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// TargetGroup defines a model for a group of targets in a level
type TargetGroup struct {
	utils.Base
	Name        string `gorm:"type:varchar(100)"`
	Description string
	SortIndex   int
	Milestone   bool
	LevelID     string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}
