package schemas

import (
	"errors"
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
			"quiz_questions": &graphql.Field{
				Type:    graphql.NewList(QuizQuestionSchema),
				Args:    FetchByIDArgument,
				Resolve: GetQuizQuestions,
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

// UpdateQuizSchema contains fields to create a new quiz
var UpdateQuizSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"title": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"target_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// QuizFromUpdateSchema is an adapter for quiz
func QuizFromUpdateSchema(p graphql.ResolveParams) models.Quiz {
	quiz := models.Quiz{}

	if title, ok := p.Args["title"]; ok {
		quiz.Title = title.(string)
	}
	if targetID, ok := p.Args["target_id"]; ok {
		quiz.TargetID = targetID.(string)
	}

	quiz.SetID(p.Args["id"].(string))

	return quiz
}

// GetQuizQuestions returns a list of quiz questions of quiz
func GetQuizQuestions(p graphql.ResolveParams) (interface{}, error) {
	quiz := p.Source.(*models.Quiz)
	if idQuery, ok := p.Args["id"].(string); ok {
		quizQuestion := models.QuizQuestion{}
		quizQuestion.SetID(idQuery)
		quizQuestion.FetchByID()
		quizQuestion.GetQuiz()
		if quizQuestion.Quiz.GetID() == quiz.GetID() {
			return models.QuizQuestionList{&quizQuestion}, nil
		}
		return nil, errors.New("Quiz question doesn't belong to quiz")
	}
	quiz.GetQuizQuestions()
	return quiz.QuizQuestions, nil
}
