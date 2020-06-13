package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

var (
	targetTableName = "targets"
)

// Target defines a model for target groups single target (equivalent to lesson)
type Target struct {
	utils.Base
	Role                   string `gorm:"type:varchar(100)"`
	Title                  string `gorm:"type:varchar(100)"`
	Description            string
	CompletionInstructions string
	ResourceURL            string `gorm:"type:varchar(255)"`
	TargetGroupID          string
	SortIndex              int
	SessionAt              *time.Time
	LinkToComplete         string `gorm:"type:varchar(255)"`
	Resubmittable          bool
	CheckList              postgres.Jsonb
	ReviewChecklist        postgres.Jsonb
}

// TableName gorm standard table name
func (u *Target) TableName() string {
	return targetTableName
}

// TargetList defines array of target objects
type TargetList []*Target

// TableName gorm standard table name
func (u *TargetList) TableName() string {
	return targetTableName
}

/**
CRUD functions
*/

// Create creates a new target record
func (u *Target) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Target by id
func (u *Target) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Targets
func (u *Target) FetchAll(ul *TargetList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given target
func (u *Target) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes target by id
func (u *Target) Delete() error {
	err := handler.Delete(u).Error
	return err
}
