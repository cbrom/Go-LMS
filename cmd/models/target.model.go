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
	TargetGroupID          string `sql:"type:uuid REFERENCES target_groups(id) ON DELETE CASCADE;" validate:"omitempty,uuid,required"`
	SortIndex              int
	SessionAt              *time.Time
	LinkToComplete         string `gorm:"type:varchar(255)"`
	Resubmittable          bool
	CheckList              postgres.Jsonb
	ReviewChecklist        postgres.Jsonb
	TargetGroup            TargetGroup       `gorm:"foreignkey:TargetGroupID"`
	TargetVersions         TargetVersionList `gorm:"foreignkey:TargetID"`
	Quizzes                QuizList          `gorm:"foreignkey:TargetID"`
}

// TableName gorm standard table name
func (t *Target) TableName() string {
	return targetTableName
}

// TargetList defines array of target objects
type TargetList []*Target

// TableName gorm standard table name
func (t *TargetList) TableName() string {
	return targetTableName
}

/**
* Relationship functions
 */

// GetTargetGroup returns target group
func (t *Target) GetTargetGroup() error {
	return handler.Model(t).Related(&t.TargetGroup).Error
}

// GetVersions returns a list of target versions
func (t *Target) GetVersions() error {
	return handler.Model(t).Related(&t.TargetVersions).Error
}

// GetQuizzes returns a list of quizzes
func (t *Target) GetQuizzes() error {
	return handler.Model(t).Related(&t.Quizzes).Error
}

/**
CRUD functions
*/

// Create creates a new target record
func (t *Target) Create() error {
	possible := handler.NewRecord(t)
	if possible {
		if err := handler.Create(t).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Target by id
func (t *Target) FetchByID() error {
	err := handler.First(t).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Targets
func (t *Target) FetchAll(tl *TargetList) error {
	err := handler.Find(tl).Error
	return err
}

// UpdateOne updates a given target
func (t *Target) UpdateOne() error {
	err := handler.Save(t).Error
	return err
}

// Delete deletes target by id
func (t *Target) Delete() error {
	err := handler.Unscoped().Delete(t).Error
	return err
}

// SoftDelete sets deleted at field
func (t *Target) SoftDelete() error {
	return handler.Delete(t).Error
}
