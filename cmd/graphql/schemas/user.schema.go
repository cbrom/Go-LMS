package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// UserSchema graphql schema of User model
var UserSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					u := p.Source.(*models.User)
					return u.GetID(), nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.Int,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"key_skills": &graphql.Field{
				Type: graphql.String,
			},
			"about": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// CreateUserSchema contains fields to create a new user
var CreateUserSchema = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"password_confirm": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"timezone": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"role": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
