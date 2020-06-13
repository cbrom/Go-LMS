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
	Title    string
	TargetID string `sql:"type:uuid" validate:"required,omitempty,uuid"`
}

// TableName gorm standard table name
func (u *Quiz) TableName() string {
	return quizTableName
}

// QuizList defines array of quiz objects
type QuizList []*Quiz

// TableName gorm standard table name
func (u *QuizList) TableName() string {
	return quizTableName
}

/**
CRUD functions
*/

// Create creates a new quiz record
func (u *Quiz) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Quiz by id
func (u *Quiz) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Quizs
func (u *Quiz) FetchAll(ul *QuizList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given quiz
func (u *Quiz) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes quiz by id
func (u *Quiz) Delete() error {
	err := handler.Delete(u).Error
	return err
}
