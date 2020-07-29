package models

import (
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/jinzhu/gorm/dialects/postgres"
)

var (
	contentBlockTableName = "content_blocks"
)

// ContentBlock defines model for blocks of a target (like markdown, image, embed, link...)
type ContentBlock struct {
	utils.Base
	BlockType       string `gorm:"type:varchar(100)"`
	Content         postgres.Jsonb
	SortIndex       int
	TargetVersionID string        `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	TargetVersion   TargetVersion `gorm:"foreignkey:TargetVersionID"`
}

// TableName gorm standard table name
func (c *ContentBlock) TableName() string {
	return contentBlockTableName
}

// ContentBlockList defines array of content block objects
type ContentBlockList []*ContentBlock

// TableName gorm standard table name
func (c *ContentBlockList) TableName() string {
	return contentBlockTableName
}

/**
* Relationship functions
 */

// GetTargetVersion returns the target version of a content block
func (c *ContentBlock) GetTargetVersion() error {
	return handler.Model(c).Related(&c.TargetVersion).Error
}

/**
CRUD functions
*/

// Create creates a new content block record
func (c *ContentBlock) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches ContentBlock by id
func (c *ContentBlock) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all ContentBlocks
func (c *ContentBlock) FetchAll(cl *ContentBlockList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given content block
func (c *ContentBlock) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes content block by id
func (c *ContentBlock) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete sets deleted at field
func (c *ContentBlock) SoftDelete() error {
	return handler.Delete(c).Error
}
