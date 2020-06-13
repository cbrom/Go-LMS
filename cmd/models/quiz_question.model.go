package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

// QuizQuestion defines a model for questions in a quiz
type QuizQuestion struct {
	utils.Base
	QuizID          string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Question        string
	Description     string `gorm:"type:varchar(100)"`
	CorrectAnswerID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}
