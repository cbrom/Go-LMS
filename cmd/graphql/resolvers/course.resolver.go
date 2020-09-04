package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateCourse creates a course
func CreateCourse(p graphql.ResolveParams) (interface{}, error) {
	course := schemas.CourseFromSchema(p)
	if err := course.Create(); err == nil {
		return course.GetID(), nil
	}

	return nil, errors.New("Unable to create course ")
}

// DeleteCourse deletes an existing course
func DeleteCourse(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		course := &models.Course{}
		course.SetID(idQuery)
		err := course.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Course id not provided")
}

// UpdateCourse updates an existing course
func UpdateCourse(p graphql.ResolveParams) (interface{}, error) {
	course := schemas.CourseFromUpdateSchema(p)
	if err := course.UpdateOne(); err == nil {
		return course.GetID(), nil
	}

	return nil, errors.New("Unable to update course")
}

// GetCourse returns a coures object for a graphql
func GetCourse(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		course := &models.Course{}
		course.SetID(idQuery)
		course.FetchByID()
		return course, nil
	}
	return nil, errors.New("Course not found")
}
