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
	LevelID     string     `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Level       *Level     `gorm:"foreignkey:LevelID"`
	Targets     TargetList `gorm:"foreignkey:TargetGroupID"`
}

// TableName gorm standard table name
func (t *TargetGroup) TableName() string {
	return targetGroupTableName
}

// TargetGroupList defines array of target group objects
type TargetGroupList []*TargetGroup

// TableName gorm standard table name
func (t *TargetGroupList) TableName() string {
	return targetGroupTableName
}

/**
CRUD functions
*/

// Create creates a new target group record
func (t *TargetGroup) Create() error {
	possible := handler.NewRecord(t)
	if possible {
		if err := handler.Create(t).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches TargetGroup by id
func (t *TargetGroup) FetchByID() error {
	err := handler.First(t).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all TargetGroups
func (t *TargetGroup) FetchAll(tl *TargetGroupList) error {
	err := handler.Find(tl).Error
	return err
}

// UpdateOne updates a given target group
func (t *TargetGroup) UpdateOne() error {
	err := handler.Save(t).Error
	return err
}

// Delete deletes target group by id
func (t *TargetGroup) Delete() error {
	err := handler.Delete(t).Error
	return err
}
