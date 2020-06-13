package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
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
