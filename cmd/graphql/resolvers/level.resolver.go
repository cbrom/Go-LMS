package resolvers

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// GetCourseLevels returns level by id with no graphql query
func GetCourseLevels(p graphql.ResolveParams) (interface{}, error) {
	courseID := p.Source.(*models.Level).GetID()
	course := &models.Course{}
	course.SetID(courseID)
	// course.GetCourses()
	return course.Levels, nil
}
