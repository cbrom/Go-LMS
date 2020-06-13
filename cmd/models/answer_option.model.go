package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// AnswerOption is a model for options of Answers for a target
type AnswerOption struct {
	utils.Base
	QuizID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Value  string
	Hint   string `gorm:"type:varchar(255)"`
}
