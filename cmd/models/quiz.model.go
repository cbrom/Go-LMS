package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	quizTableName = "quizzes"
)

// Quiz defines a model for target quizes
type Quiz struct {
	utils.Base
	Title         string
	TargetID      string           `sql:"type:uuid" validate:"required,omitempty,uuid"`
	Target        *Target          `gorm:"foreignkey:TargetID"`
	QuizQuestions QuizQuestionList `gorm:"foreignkey:QuizID"`
}

// TableName gorm standard table name
func (q *Quiz) TableName() string {
	return quizTableName
}

// QuizList defines array of quiz objects
type QuizList []*Quiz

// TableName gorm standard table name
func (q *QuizList) TableName() string {
	return quizTableName
}

/**
CRUD functions
*/

// Create creates a new quiz record
func (q *Quiz) Create() error {
	possible := handler.NewRecord(q)
	if possible {
		if err := handler.Create(q).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Quiz by id
func (q *Quiz) FetchByID() error {
	err := handler.First(q).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Quizs
func (q *Quiz) FetchAll(ql *QuizList) error {
	err := handler.Find(ql).Error
	return err
}

// UpdateOne updates a given quiz
func (q *Quiz) UpdateOne() error {
	err := handler.Save(q).Error
	return err
}

// Delete deletes quiz by id
func (q *Quiz) Delete() error {
	err := handler.Delete(q).Error
	return err
}
