package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	targetGroupTableName = "target_groups"
)

// TargetGroup defines a model for a group of targets in a level
type TargetGroup struct {
	utils.Base
	Name        string `gorm:"type:varchar(100)"`
	Description string
	SortIndex   int
	Milestone   bool
	LevelID     string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}

// TableName gorm standard table name
func (u *TargetGroup) TableName() string {
	return targetGroupTableName
}

// TargetGroupList defines array of target group objects
type TargetGroupList []*TargetGroup

// TableName gorm standard table name
func (u *TargetGroupList) TableName() string {
	return targetGroupTableName
}

/**
CRUD functions
*/

// Create creates a new target group record
func (u *TargetGroup) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches TargetGroup by id
func (u *TargetGroup) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all TargetGroups
func (u *TargetGroup) FetchAll(ul *TargetGroupList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given target group
func (u *TargetGroup) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes target group by id
func (u *TargetGroup) Delete() error {
	err := handler.Delete(u).Error
	return err
}
