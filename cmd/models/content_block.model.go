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
	TargetVersionID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}

// TableName gorm standard table name
func (u *ContentBlock) TableName() string {
	return contentBlockTableName
}

// ContentBlockList defines array of content block objects
type ContentBlockList []*ContentBlock

// TableName gorm standard table name
func (u *ContentBlockList) TableName() string {
	return contentBlockTableName
}

/**
CRUD functions
*/

// Create creates a new content block record
func (u *ContentBlock) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches ContentBlock by id
func (u *ContentBlock) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all ContentBlocks
func (u *ContentBlock) FetchAll(ul *ContentBlockList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given content block
func (u *ContentBlock) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes content block by id
func (u *ContentBlock) Delete() error {
	err := handler.Delete(u).Error
	return err
}
