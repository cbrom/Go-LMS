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
	Name        string `gorm:"type:varchar(100)"`
	CourseID    string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Description string
	Number      int
	UnlockOn    *time.Time
}

// TableName gorm standard table name
func (u *Level) TableName() string {
	return levelTableName
}

// LevelList defines array of level objects
type LevelList []*Level

// TableName gorm standard table name
func (u *LevelList) TableName() string {
	return levelTableName
}

/**
CRUD functions
*/

// Create creates a new level record
func (u *Level) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Level by id
func (u *Level) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Levels
func (u *Level) FetchAll(ul *LevelList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given level
func (u *Level) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes level by id
func (u *Level) Delete() error {
	err := handler.Delete(u).Error
	return err
}
