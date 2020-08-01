package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// QuizSchema graphql schema of Quiz model
var QuizSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Quiz",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					q := p.Source.(*models.Quiz)
					return q.GetID(), nil
				},
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// CreateQuizSchema contains fields to create a new quiz
var CreateQuizSchema = graphql.FieldConfigArgument{
	"title": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"target_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// QuizFromSchema is an adapter for quiz
func QuizFromSchema(p graphql.ResolveParams) models.Quiz {
	quiz := models.Quiz{
		Title:    p.Args["title"].(string),
		TargetID: p.Args["target_id"].(string),
	}

	return quiz
}
