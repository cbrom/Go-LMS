package models

import (
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// ContentBlock defines model for blocks of a target (like markdown, image, embed, link...)
type ContentBlock struct {
	utils.Base
	BlockType       string `gorm:"type:varchar(100)"`
	Content         postgres.Jsonb
	SortIndex       int
	TargetVersionID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}
