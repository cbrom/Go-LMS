package schemas

import (
	"errors"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// QuizQuestionSchema graphql schema of quiz question model
var QuizQuestionSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "QuizQuestion",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					qq := p.Source.(*models.QuizQuestion)
					return qq.GetID(), nil
				},
			},
			"question": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"answer_options": &graphql.Field{
				Type:    graphql.NewList(AnswerOptionSchema),
				Args:    FetchByIDArgument,
				Resolve: GetAnswerOptions,
			},
		},
	})

// CreateQuizQuestionSchema contains fields to create a new aquiz questionn
var CreateQuizQuestionSchema = graphql.FieldConfigArgument{
	"quiz_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"question": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"correct_answer_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// QuizQuestionFromSchema is an adapter for quiz question
func QuizQuestionFromSchema(p graphql.ResolveParams) models.QuizQuestion {
	quizQuestion := models.QuizQuestion{
		QuizID:      p.Args["quiz_id"].(string),
		Question:    p.Args["question"].(string),
		Description: p.Args["description"].(string),
	}

	return quizQuestion
}

// UpdateQuizQuestionSchema contains fields to updates a quiz questionn
var UpdateQuizQuestionSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"quiz_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"question": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"correct_answer_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// QuizQuestionFromUpdateSchema is an adapter for quiz question
func QuizQuestionFromUpdateSchema(p graphql.ResolveParams) models.QuizQuestion {
	quizQuestion := models.QuizQuestion{}

	if quizID, ok := p.Args["quiz_id"]; ok {
		quizQuestion.QuizID = quizID.(string)
	}
	if question, ok := p.Args["question"]; ok {
		quizQuestion.Question = question.(string)
	}
	if description, ok := p.Args["description"]; ok {
		quizQuestion.Description = description.(string)
	}

	quizQuestion.SetID(p.Args["id"].(string))
	return quizQuestion
}

// GetAnswerOptions returns a list of answer options for quiz question
func GetAnswerOptions(p graphql.ResolveParams) (interface{}, error) {
	quizQuestion := p.Source.(*models.QuizQuestion)
	if idQuery, ok := p.Args["id"].(string); ok {
		answerOption := models.AnswerOption{}
		answerOption.SetID(idQuery)
		answerOption.FetchByID()
		answerOption.GetQuestion()
		if answerOption.QuizQuestion.GetID() == quizQuestion.GetID() {
			return models.AnswerOptionList{&answerOption}, nil
		}
		return nil, errors.New("Answer option doesn't belong to quiz question")
	}

	quizQuestion.GetAnswerOptions()
	return quizQuestion.Answers, nil
}
