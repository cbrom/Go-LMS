package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateQuiz creates a new quiz
func CreateQuiz(p graphql.ResolveParams) (interface{}, error) {
	quiz := schemas.QuizFromSchema(p)
	if err := quiz.Create(); err == nil {
		return quiz.GetID(), nil
	}

	return nil, errors.New("Unable to create quiz")
}

// CreateQuizQuestion creates a new quizQuestion
func CreateQuizQuestion(p graphql.ResolveParams) (interface{}, error) {
	quizQuestion := schemas.QuizQuestionFromSchema(p)
	if err := quizQuestion.Create(); err == nil {
		return quizQuestion.GetID(), nil
	}

	return nil, errors.New("Unable to create quizQuestion")
}

// CreateAnswerOption creates a new answerOption
func CreateAnswerOption(p graphql.ResolveParams) (interface{}, error) {
	answerOption := schemas.AnswerOptionFromSchema(p)
	if err := answerOption.Create(); err == nil {
		return answerOption.GetID(), nil
	}

	return nil, errors.New("Unable to create answerOption")
}
