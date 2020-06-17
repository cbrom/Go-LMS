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
	Description     string           `gorm:"type:varchar(100)"`
	CorrectAnswerID string           `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Quiz            *Quiz            `gorm:"foreignkey:QuizID"`
	Answer          *AnswerOption    `gorm:"foreignkey:CorrectAnswerID"`
	Answers         AnswerOptionList `gorm:"foreign:QuizQuestionID"`
}

// TableName gorm standard table name
func (i *QuizQuestion) TableName() string {
	return quizQuestionTableName
}

// QuizQuestionList defines array of quiz question objects
type QuizQuestionList []*QuizQuestion

// TableName gorm standard table name
func (i *QuizQuestionList) TableName() string {
	return quizQuestionTableName
}

/**
CRUD functions
*/

// Create creates a new quiz question record
func (i *QuizQuestion) Create() error {
	possible := handler.NewRecord(i)
	if possible {
		if err := handler.Create(i).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches QuizQuestion by id
func (i *QuizQuestion) FetchByID() error {
	err := handler.First(i).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all QuizQuestions
func (i *QuizQuestion) FetchAll(il *QuizQuestionList) error {
	err := handler.Find(il).Error
	return err
}

// UpdateOne updates a given quiz question
func (i *QuizQuestion) UpdateOne() error {
	err := handler.Save(i).Error
	return err
}

// Delete deletes quiz question by id
func (i *QuizQuestion) Delete() error {
	err := handler.Delete(i).Error
	return err
}
