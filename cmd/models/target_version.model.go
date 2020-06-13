package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// TargetVersion defines a model for a specific version of a target
type TargetVersion struct {
	utils.Base
	TargetID    string `sql:"type:uuid" validate:"omitempty,required,uuid"`
	VersionName string `gorm:"type:varchar(100)"`
}
