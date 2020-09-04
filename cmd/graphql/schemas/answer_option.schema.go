package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// AnswerOptionSchema graphql schema of answer options model
var AnswerOptionSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AnswerOption",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ao := p.Source.(*models.AnswerOption)
					return ao.GetID(), nil
				},
			},
			"value": &graphql.Field{
				Type: graphql.String,
			},
			"hint": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// CreateAnswerOptionSchema contains fields to create a new answer option
var CreateAnswerOptionSchema = graphql.FieldConfigArgument{
	"quiz_question_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"value": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"hint": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// AnswerOptionFromSchema is an adapter for answer option
func AnswerOptionFromSchema(p graphql.ResolveParams) models.AnswerOption {
	answerOption := models.AnswerOption{
		QuizQuestionID: p.Args["quiz_question_id"].(string),
		Value:          p.Args["value"].(string),
		Hint:           p.Args["hint"].(string),
	}
	return answerOption
}


// UpdateAnswerOptionSchema contains fields to update an answer option
var UpdateAnswerOptionSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"quiz_question_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"value": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"hint": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// AnswerOptionFromUpdateSchema is an adapter for update answer option schema
func AnswerOptionFromUpdateSchema(p graphql.ResolveParams) models.AnswerOption {
	answerOption := models.AnswerOption{
		QuizQuestionID: p.Args["quiz_question_id"].(string),
		Value:          p.Args["value"].(string),
		Hint:           p.Args["hint"].(string),
	}
	answerOption.SetID(p.Args["id"].(string))
	return answerOption
}