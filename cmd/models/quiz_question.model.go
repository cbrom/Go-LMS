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
	CorrectAnswerID string
	Quiz            Quiz             `gorm:"foreignkey:QuizID"`
	Answer          *AnswerOption    `gorm:"foreignkey:CorrectAnswerID"`
	Answers         AnswerOptionList `gorm:"foreign:QuizQuestionID"`
}

// TableName gorm standard table name
func (q *QuizQuestion) TableName() string {
	return quizQuestionTableName
}

// QuizQuestionList defines array of quiz question objects
type QuizQuestionList []*QuizQuestion

// TableName gorm standard table name
func (q *QuizQuestionList) TableName() string {
	return quizQuestionTableName
}

/**
Relationship functions
*/

// GetAnswerOptions returns answer options of a question
func (q *QuizQuestion) GetAnswerOptions() error {
	return handler.Model(q).Related(&q.Answers).Error
}

// GetAnswer returns the answer of a question
func (q *QuizQuestion) GetAnswer() error {
	answer := AnswerOption{}
	err := handler.Model(q).Related(&q.Answer).Error
	q.Answer = &answer
	return err
}

// GetQuiz returns the quiz of this question
func (q *QuizQuestion) GetQuiz() error {
	return handler.Model(q).Related(&q.Quiz).Error
}

/**
CRUD functions
*/

// Create creates a new quiz question record
func (q *QuizQuestion) Create() error {
	possible := handler.NewRecord(q)
	if possible {
		if err := handler.Create(q).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches QuizQuestion by id
func (q *QuizQuestion) FetchByID() error {
	err := handler.First(q).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all QuizQuestions
func (q *QuizQuestion) FetchAll(ql *QuizQuestionList) error {
	err := handler.Find(ql).Error
	return err
}

// UpdateOne updates a given quiz question
func (q *QuizQuestion) UpdateOne() error {
	err := handler.Save(q).Error
	return err
}

// Delete deletes quiz question by id
func (q *QuizQuestion) Delete() error {
	err := handler.Unscoped().Delete(q).Error
	return err
}

// SoftDelete sets deleted at field
func (q *QuizQuestion) SoftDelete() error {
	return handler.Delete(q).Error
}
