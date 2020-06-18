package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// LevelSchema graphql schema of Level model
var LevelSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Level",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					l := p.Source.(*models.Level)
					return l.GetID(), nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"number": &graphql.Field{
				Type: graphql.Int,
			},
			"unlock_on": &graphql.Field{
				Type: graphql.String,
			},
			"course": &graphql.Field{
				Type: CourseSchema,
			},
		},
	})

// CreateLevelSchema contains fields to create a new user
var CreateLevelSchema = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"number": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"unlock_on": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
