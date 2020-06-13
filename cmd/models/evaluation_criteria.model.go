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
}

// TableName gorm standard table name
func (u *EvaluationCriteria) TableName() string {
	return evaluationCriteriaTableName
}

// EvaluationCriteriaList defines array of evaluation criteria objects
type EvaluationCriteriaList []*EvaluationCriteria

// TableName gorm standard table name
func (u *EvaluationCriteriaList) TableName() string {
	return evaluationCriteriaTableName
}

/**
CRUD functions
*/

// Create creates a new evaluation criteria record
func (u *EvaluationCriteria) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches EvaluationCriteria by id
func (u *EvaluationCriteria) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all EvaluationCriterias
func (u *EvaluationCriteria) FetchAll(ul *EvaluationCriteriaList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given evaluation criteria
func (u *EvaluationCriteria) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes evaluation criteria by id
func (u *EvaluationCriteria) Delete() error {
	err := handler.Delete(u).Error
	return err
}
