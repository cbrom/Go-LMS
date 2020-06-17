package models

import (
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/jinzhu/gorm/dialects/postgres"
)

var (
	evaluationCriteriaTableName = "evaluation_criterias"
)

// EvaluationCriteria defines a model for course evaluation criteria
type EvaluationCriteria struct {
	utils.Base
	Name        string `gorm:"type:varchar(100)"`
	CourseID    string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	MaxGrade    uint
	PassGrade   uint
	GradeLabels postgres.Jsonb
	Course      *Course `gorm:"foreignkey:CourseID"`
}

// TableName gorm standard table name
func (e *EvaluationCriteria) TableName() string {
	return evaluationCriteriaTableName
}

// EvaluationCriteriaList defines array of evaluation criteria objects
type EvaluationCriteriaList []*EvaluationCriteria

// TableName gorm standard table name
func (e *EvaluationCriteriaList) TableName() string {
	return evaluationCriteriaTableName
}

/**
CRUD functions
*/

// Create creates a new evaluation criteria record
func (e *EvaluationCriteria) Create() error {
	possible := handler.NewRecord(e)
	if possible {
		if err := handler.Create(e).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches EvaluationCriteria by id
func (e *EvaluationCriteria) FetchByID() error {
	err := handler.First(e).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all EvaluationCriterias
func (e *EvaluationCriteria) FetchAll(el *EvaluationCriteriaList) error {
	err := handler.Find(el).Error
	return err
}

// UpdateOne updates a given evaluation criteria
func (e *EvaluationCriteria) UpdateOne() error {
	err := handler.Save(e).Error
	return err
}

// Delete deletes evaluation criteria by id
func (e *EvaluationCriteria) Delete() error {
	err := handler.Delete(e).Error
	return err
}
