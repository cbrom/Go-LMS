package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	answerOptionTableName = "answer_options"
)

// AnswerOption is a model for options of Answers for a target
type AnswerOption struct {
	utils.Base
	QuizID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Value  string
	Hint   string `gorm:"type:varchar(255)"`
}

// TableName gorm standard table name
func (u *AnswerOption) TableName() string {
	return answerOptionTableName
}

// AnswerOptionList defines array of user objects
type AnswerOptionList []*AnswerOption

// TableName gorm standard table name
func (u *AnswerOptionList) TableName() string {
	return answerOptionTableName
}

/**
CRUD functions
*/

// Create creates a new user record
func (u *AnswerOption) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches AnswerOption by id
func (u *AnswerOption) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all AnswerOptions
func (u *AnswerOption) FetchAll(ul *AnswerOptionList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given user
func (u *AnswerOption) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes user by id
func (u *AnswerOption) Delete() error {
	err := handler.Delete(u).Error
	return err
}
