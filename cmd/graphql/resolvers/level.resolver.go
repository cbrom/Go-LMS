package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateLevel creates a new level
func CreateLevel(p graphql.ResolveParams) (interface{}, error) {
	level := schemas.CreateLevelFromSchema(p)
	if err := level.Create(); err == nil {
		return level.GetID(), nil
	}
	return nil, errors.New("Unable to create level")
}

// GetCourseLevels returns level by id with no graphql query
func GetCourseLevels(p graphql.ResolveParams) (interface{}, error) {
	courseID := p.Source.(*models.Level).GetID()
	course := &models.Course{}
	course.SetID(courseID)
	// course.GetCourses()
	return course.Levels, nil
}
