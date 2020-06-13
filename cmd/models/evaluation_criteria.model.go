package models

import (
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// EvaluationCriteria defines a model for course evaluation criteria
type EvaluationCriteria struct {
	utils.Base
	Name        string `gorm:"type:varchar(100)"`
	CourseID    string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	MaxGrade    uint
	PassGrade   uint
	GradeLabels postgres.Jsonb
}
