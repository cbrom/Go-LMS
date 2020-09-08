package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// CourseAuthorSchema graphql schema of course author model
var CourseAuthorSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CourseAuthor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ca := p.Source.(*models.CourseAuthor)
					return ca.GetID(), nil
				},
			},
		},
	})

// CreateCourseAuthorSchema contains fields to create a new course author
var CreateCourseAuthorSchema = graphql.FieldConfigArgument{
	"user_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// CourseAuthorFromSchema is an adapter for course author
func CourseAuthorFromSchema(p graphql.ResolveParams) models.CourseAuthor {
	courseAuthor := models.CourseAuthor{
		UserID:   p.Args["user_id"].(string),
		CourseID: p.Args["course_id"].(string),
	}

	return courseAuthor
}

// UpdateCourseAuthorSchema contains fields to update a course author
var UpdateCourseAuthorSchema = graphql.FieldConfigArgument{
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

// CourseAuthorFromUpdateSchema is an adapter for course author
func CourseAuthorFromUpdateSchema(p graphql.ResolveParams) models.CourseAuthor {
	courseAuthor := models.CourseAuthor{}
	if userID, ok := p.Args["user_id"]; ok {
		courseAuthor.UserID = userID.(string)
	}
	if courseID, ok := p.Args["course_id"]; ok {
		courseAuthor.CourseID = courseID.(string)
	}
	courseAuthor.SetID(p.Args["id"].(string))

	return courseAuthor
}
