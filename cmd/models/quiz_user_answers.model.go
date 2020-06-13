package models

import "go-lms-of-pupilfirst/pkg/utils"

var (
	quizQuestionUserAnswerTableName = "quiz_question_user_answers"
)

// QuizUserAnswer defines a model for questions in a quiz
type QuizUserAnswer struct {
	utils.Base
	QuestionID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	AnswerID   string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	UserID     string `sql:"type:uuid" validate:"omitempty,uuid,required"`
}

// TableName gorm standard table name
func (u *QuizUserAnswer) TableName() string {
	return quizQuestionUserAnswerTableName
}

// QuizUserAnswerList defines array of quiz user answer objects
type QuizUserAnswerList []*QuizUserAnswer

// TableName gorm standard table name
func (u *QuizUserAnswerList) TableName() string {
	return quizQuestionUserAnswerTableName
}

/**
CRUD functions
*/

// Create creates a new quiz user answer record
func (u *QuizUserAnswer) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches QuizUserAnswer by id
func (u *QuizUserAnswer) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all QuizUserAnswers
func (u *QuizUserAnswer) FetchAll(ul *QuizUserAnswerList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given quiz user answer
func (u *QuizUserAnswer) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes quiz user answer by id
func (u *QuizUserAnswer) Delete() error {
	err := handler.Delete(u).Error
	return err
}
