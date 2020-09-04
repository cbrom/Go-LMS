package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

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

// UpdateQuiz updates an existing quiz
func UpdateQuiz(p graphql.ResolveParams) (interface{}, error) {
	quiz := schemas.QuizFromUpdateSchema(p)
	if err := quiz.UpdateOne(); err == nil {
		return quiz.GetID(), nil
	}

	return nil, errors.New("Unable to update quiz ")
}

// DeleteQuiz deletes an existing quiz
func DeleteQuiz(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		quiz := &models.Quiz{}
		quiz.SetID(idQuery)
		err := quiz.SoftDelete() 
		return nil, err
	}
	return nil, errors.New("Quiz id not provided")
}

// CreateQuizQuestion creates a new quizQuestion
func CreateQuizQuestion(p graphql.ResolveParams) (interface{}, error) {
	quizQuestion := schemas.QuizQuestionFromSchema(p)
	if err := quizQuestion.Create(); err == nil {
		return quizQuestion.GetID(), nil
	}

	return nil, errors.New("Unable to create quizQuestion")
}

// UpdateQuizQuestion updates an existing quiz question
func UpdateQuizQuestion(p graphql.ResolveParams) (interface{}, error) {
	quizQuestion := schemas.QuizQuestionFromUpdateSchema(p)
	if err := quizQuestion.UpdateOne(); err == nil {
		return quizQuestion.GetID(), nil
	}

	return nil, errors.New("Unable to update quiz question ")
}

// DeleteQuizQuestion deletes an existing quiz question
func DeleteQuizQuestion(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		quizQuestion := &models.QuizQuestion{}
		quizQuestion.SetID(idQuery)
		err := quizQuestion.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Quiz question id not provided")
}

// CreateAnswerOption creates a new answerOption
func CreateAnswerOption(p graphql.ResolveParams) (interface{}, error) {
	answerOption := schemas.AnswerOptionFromSchema(p)
	if err := answerOption.Create(); err == nil {
		return answerOption.GetID(), nil
	}

	return nil, errors.New("Unable to create answerOption")
}

// UpdateAnswerOption updates an existing answer option
func UpdateAnswerOption(p graphql.ResolveParams) (interface{}, error) {
	answerOption := schemas.AnswerOptionFromUpdateSchema(p)
	if err := answerOption.UpdateOne(); err == nil {
		return answerOption.GetID(), nil
	}

	return nil, errors.New("Unable to update answer option ")
}

// DeleteAnswerOption deletes an existing answer option
func DeleteAnswerOption(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		answerOption := &models.AnswerOption{}
		answerOption.SetID(idQuery)
		err := answerOption.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Answer option id not provided")
}
