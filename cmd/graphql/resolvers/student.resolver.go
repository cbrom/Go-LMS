package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// GetStudent returns a student
func GetStudent(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		user := models.User{}
		user.SetID(idQuery)
		user.FetchByID()
		student := schemas.Student{
			ID:       user.GetID(),
			UserInfo: user,
		}
		return student, nil
	}
	return nil, errors.New("Unable to find student")
}

// EnrollStudentInCourse enrolles a student into a course
func EnrollStudentInCourse(p graphql.ResolveParams) (interface{}, error) {
	studentCourse := schemas.StudentCourseFromSchema(p)
	if err := studentCourse.Create(); err == nil {
		return studentCourse.GetID(), nil
	}

	return nil, errors.New("Unable enroll student in a coures")
}
