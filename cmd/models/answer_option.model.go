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
	QuizQuestionID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Value          string
	Hint           string `gorm:"type:varchar(255)"`

	QuizQuestion QuizQuestion `gorm:"foreignkey:QuizQuestionID"`
}

// TableName gorm standard table name
func (a *AnswerOption) TableName() string {
	return answerOptionTableName
}

// AnswerOptionList defines array of answer options objects
type AnswerOptionList []*AnswerOption

// TableName gorm standard table name
func (a *AnswerOptionList) TableName() string {
	return answerOptionTableName
}

/**
Relationship functions
*/

// GetQuestion returns the QuizQuestion of this answer
func (a *AnswerOption) GetQuestion() error {
	return handler.Model(a).Related(&a.QuizQuestion).Error
}

/**
CRUD functions
*/

// Create creates a new answer options record
func (a *AnswerOption) Create() error {
	possible := handler.NewRecord(a)
	if possible {
		if err := handler.Create(a).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches AnswerOption by id
func (a *AnswerOption) FetchByID() error {
	err := handler.First(a).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all AnswerOptions
func (a *AnswerOption) FetchAll(al *AnswerOptionList) error {
	err := handler.Find(al).Error
	return err
}

// UpdateOne updates a given answer options
func (a *AnswerOption) UpdateOne() error {
	err := handler.Save(a).Error
	return err
}

// Delete deletes answer options by id
func (a *AnswerOption) Delete() error {
	err := handler.Unscoped().Delete(a).Error
	return err
}

// SoftDelete sets deleted at field
func (a *AnswerOption) SoftDelete() error {
	return handler.Delete(a).Error
}
