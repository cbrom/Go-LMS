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
	TargetID    string `sql:"type:uuid" validate:"omitempty,required,uuid"`
	VersionName string `gorm:"type:varchar(100)"`
}

// TableName gorm standard table name
func (u *TargetVersion) TableName() string {
	return targetVersionTableName
}

// TargetVersionList defines array of target version objects
type TargetVersionList []*TargetVersion

// TableName gorm standard table name
func (u *TargetVersionList) TableName() string {
	return targetVersionTableName
}

/**
CRUD functions
*/

// Create creates a new target version record
func (u *TargetVersion) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches TargetVersion by id
func (u *TargetVersion) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all TargetVersions
func (u *TargetVersion) FetchAll(ul *TargetVersionList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given target version
func (u *TargetVersion) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes target version by id
func (u *TargetVersion) Delete() error {
	err := handler.Delete(u).Error
	return err
}
