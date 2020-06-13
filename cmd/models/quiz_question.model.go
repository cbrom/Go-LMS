package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	quizQuestionTableName = "quiz_questions"
)

// QuizQuestion defines a model for questions in a quiz
type QuizQuestion struct {
	utils.Base
	QuizID          string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Question        string
	Description     string `gorm:"type:varchar(100)"`
	CorrectAnswerID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}

// TableName gorm standard table name
func (u *QuizQuestion) TableName() string {
	return quizQuestionTableName
}

// QuizQuestionList defines array of quiz question objects
type QuizQuestionList []*QuizQuestion

// TableName gorm standard table name
func (u *QuizQuestionList) TableName() string {
	return quizQuestionTableName
}

/**
CRUD functions
*/

// Create creates a new quiz question record
func (u *QuizQuestion) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches QuizQuestion by id
func (u *QuizQuestion) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all QuizQuestions
func (u *QuizQuestion) FetchAll(ul *QuizQuestionList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given quiz question
func (u *QuizQuestion) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes quiz question by id
func (u *QuizQuestion) Delete() error {
	err := handler.Delete(u).Error
	return err
}
