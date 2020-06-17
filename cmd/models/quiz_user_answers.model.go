package models

import "go-lms-of-pupilfirst/pkg/utils"

var (
	quizQuestionUserAnswerTableName = "quiz_question_user_answers"
)

// QuizUserAnswer defines a model for questions in a quiz
type QuizUserAnswer struct {
	utils.Base
	QuestionID string        `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	AnswerID   string        `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	UserID     string        `sql:"type:uuid" validate:"omitempty,uuid,required"`
	Question   *QuizQuestion `gorm:"foreignkey:QuestionID"`
	Answer     *AnswerOption `gorm:"foreignkey:AnswerID"`
	User       *User         `gorm:"foreignkey:UserID"`
}

// TableName gorm standard table name
func (q *QuizUserAnswer) TableName() string {
	return quizQuestionUserAnswerTableName
}

// QuizUserAnswerList defines array of quiz user answer objects
type QuizUserAnswerList []*QuizUserAnswer

// TableName gorm standard table name
func (q *QuizUserAnswerList) TableName() string {
	return quizQuestionUserAnswerTableName
}

/**
CRUD functions
*/

// Create creates a new quiz user answer record
func (q *QuizUserAnswer) Create() error {
	possible := handler.NewRecord(q)
	if possible {
		if err := handler.Create(q).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches QuizUserAnswer by id
func (q *QuizUserAnswer) FetchByID() error {
	err := handler.First(q).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all QuizUserAnswers
func (q *QuizUserAnswer) FetchAll(ql *QuizUserAnswerList) error {
	err := handler.Find(ql).Error
	return err
}

// UpdateOne updates a given quiz user answer
func (q *QuizUserAnswer) UpdateOne() error {
	err := handler.Save(q).Error
	return err
}

// Delete deletes quiz user answer by id
func (q *QuizUserAnswer) Delete() error {
	err := handler.Delete(q).Error
	return err
}
