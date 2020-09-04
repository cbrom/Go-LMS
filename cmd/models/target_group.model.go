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
	LevelID     string     `sql:"type:uuid REFERENCES levels(id) ON DELETE CASCADE;" validate:"omitempty,uuid,required"`
	Level       Level      `gorm:"foreignkey:LevelID"`
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
* Relationship functions
 */

// GetLevel returns the level of the target group
func (t *TargetGroup) GetLevel() error {
	return handler.Model(t).Related(&t.Level).Error
}

// GetTargets returns the list of targets
func (t *TargetGroup) GetTargets() error {
	return handler.Model(t).Related(&t.Targets).Error
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
	err := handler.Unscoped().Delete(t).Error
	return err
}

// SoftDelete sets deleted at field
func (t *TargetGroup) SoftDelete() error {
	return handler.Delete(t).Error
}
