package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// StudentSchema graphql schema of student model
var StudentSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_info": &graphql.Field{
				Type:    UserSchema,
				Resolve: GetStudentUser,
			},
			"courses": &graphql.Field{
				Type:    graphql.NewList(CourseSchema),
				Resolve: GetStudentCourses,
			},
			"certificates": &graphql.Field{
				Args:    FetchByIDArgument,
				Type:    graphql.NewList(CertificateSchema),
				Resolve: GetStudentCertificates,
			},
		},
	})

// Student is a schema for a student
type Student struct {
	ID       string
	UserInfo models.User
}

// StudentCourseSchema graphql schema of student course model
var StudentCourseSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StudentCourse",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					s := p.Source.(*models.StudentCourse)
					return s.GetID(), nil
				},
			},
		},
	})

// CreateStudentCourseSchema contains fields to create a new student course
var CreateStudentCourseSchema = graphql.FieldConfigArgument{
	"user_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// StudentCourseFromSchema is an adapter for student course
func StudentCourseFromSchema(p graphql.ResolveParams) models.StudentCourse {
	studentCourse := models.StudentCourse{
		UserID:   p.Args["user_id"].(string),
		CourseID: p.Args["course_id"].(string),
	}

	return studentCourse
}

// UpdateStudentCourseSchema contains fields to update a student course
var UpdateStudentCourseSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"user_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// StudentCourseFromUpdateSchema is an adapter for student course
func StudentCourseFromUpdateSchema(p graphql.ResolveParams) models.StudentCourse {
	studentCourse := models.StudentCourse{}

	if userID, ok := p.Args["user_id"]; ok {
		studentCourse.UserID = userID.(string)
	}
	if courseID, ok := p.Args["course_id"]; ok {
		studentCourse.CourseID = courseID.(string)
	}

	studentCourse.SetID(p.Args["id"].(string))
	return studentCourse
}

// GetStudentUser  returns user info of a student
func GetStudentUser(p graphql.ResolveParams) (interface{}, error) {
	student := p.Source.(Student)

	return student.UserInfo, nil
}

// GetStudentCertificates fetches user's certificates
func GetStudentCertificates(p graphql.ResolveParams) (interface{}, error) {
	student := p.Source.(Student)
	student.UserInfo.GetStudentCertificates()
	return student.UserInfo.StudentCertificates, nil
}

// GetStudentCourses returns a list of student courses
func GetStudentCourses(p graphql.ResolveParams) (interface{}, error) {
	student := p.Source.(Student)
	student.UserInfo.GetCourses()
	return student.UserInfo.AllCourses, nil
}
