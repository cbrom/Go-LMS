package models

import "go-lms-of-pupilfirst/pkg/utils"

// QuizUserAnswer defines a model for questions in a quiz
type QuizUserAnswer struct {
	utils.Base
	QuestionID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	AnswerID   string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	UserID     string `sql:"type:uuid" validate:"omitempty,uuid,required"`
}
