package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	targetVersionTableName = "target_versions"
)

// TargetVersion defines a model for a specific version of a target
type TargetVersion struct {
	utils.Base
	TargetID      string           `sql:"type:uuid" validate:"omitempty,required,uuid"`
	VersionName   string           `gorm:"type:varchar(100)"`
	Target        Target           `gorm:"foreignkey:TargetID"`
	ContentBlocks ContentBlockList `gorm:"foreignkey:TargetVersion"`
}

// TableName gorm standard table name
func (t *TargetVersion) TableName() string {
	return targetVersionTableName
}

// TargetVersionList defines array of target version objects
type TargetVersionList []*TargetVersion

// TableName gorm standard table name
func (t *TargetVersionList) TableName() string {
	return targetVersionTableName
}

/**
* Relationship functions
 */

// GetTarget returns target of this version
func (t *TargetVersion) GetTarget() error {
	return handler.Model(t).Related(&t.Target).Error
}

// GetContentBlocks returns content blocks of this version
func (t *TargetVersion) GetContentBlocks() error {
	return handler.Model(t).Related(&t.ContentBlocks).Error
}

/**
CRUD functions
*/

// Create creates a new target version record
func (t *TargetVersion) Create() error {
	possible := handler.NewRecord(t)
	if possible {
		if err := handler.Create(t).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches TargetVersion by id
func (t *TargetVersion) FetchByID() error {
	err := handler.First(t).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all TargetVersions
func (t *TargetVersion) FetchAll(tl *TargetVersionList) error {
	err := handler.Find(tl).Error
	return err
}

// UpdateOne updates a given target version
func (t *TargetVersion) UpdateOne() error {
	err := handler.Save(t).Error
	return err
}

// Delete deletes target version by id
func (t *TargetVersion) Delete() error {
	err := handler.Unscoped().Delete(t).Error
	return err
}

// SoftDelete sets deleted at field
func (t *TargetVersion) SoftDelete() error {
	return handler.Delete(t).Error
}
