package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"
)

var (
	levelTableName = "levels"
)

// Level defines a model for course levels
type Level struct {
	utils.Base
	Name         string `gorm:"type:varchar(100)"`
	CourseID     string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Description  string
	Number       int
	UnlockOn     *time.Time
	Course       *Course         `gorm:"foreignkey:CourseID"`
	TargetGroups TargetGroupList `gorm:"foreignkey:TargetGroupID"`
}

// TableName gorm standard table name
func (l *Level) TableName() string {
	return levelTableName
}

// LevelList defines array of level objects
type LevelList []*Level

// TableName gorm standard table name
func (l *LevelList) TableName() string {
	return levelTableName
}

/**
CRUD functions
*/

// Create creates a new level record
func (l *Level) Create() error {
	possible := handler.NewRecord(l)
	if possible {
		if err := handler.Create(l).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Level by id
func (l *Level) FetchByID() error {
	err := handler.First(l).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Levels
func (l *Level) FetchAll(ll *LevelList) error {
	err := handler.Find(ll).Error
	return err
}

// UpdateOne updates a given level
func (l *Level) UpdateOne() error {
	err := handler.Save(l).Error
	return err
}

// Delete deletes level by id
func (l *Level) Delete() error {
	err := handler.Delete(l).Error
	return err
}
