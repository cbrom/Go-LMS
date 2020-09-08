package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/graphql-go/graphql"
)

// EvaluationCriteriaSchema graphql schema of evaluation criteria model
var EvaluationCriteriaSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EvaluationCriteria",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ec := p.Source.(*models.EvaluationCriteria)
					return ec.GetID(), nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"max_grade": &graphql.Field{
				Type: graphql.Int,
			},
			"pass_grade": &graphql.Field{
				Type: graphql.Int,
			},
			"grade_labels": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// CreateEvaluationCriteriaSchema contains fields to create a new evaluation criteria
var CreateEvaluationCriteriaSchema = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"max_grade": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"pass_grade": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"grade_labels": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// EvaluationCriteriaFromSchema is an adapter for evaluation criteria
func EvaluationCriteriaFromSchema(p graphql.ResolveParams) models.EvaluationCriteria {
	evaluationCriteria := models.EvaluationCriteria{
		Name:        p.Args["name"].(string),
		CourseID:    p.Args["course_id"].(string),
		MaxGrade:    uint(p.Args["max_grade"].(int)),
		PassGrade:   uint(p.Args["pass_grade"].(int)),
		GradeLabels: utils.ConvertStringToJsonb(p.Args["grade_labels"].(string)),
	}

	return evaluationCriteria
}

// UpdateEvaluationCriteriaSchema contains fields to update an evaluation criteria
var UpdateEvaluationCriteriaSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"max_grade": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"pass_grade": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"grade_labels": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// EvaluationCriteriaFromUpdateSchema is an adapter for evaluation criteria
func EvaluationCriteriaFromUpdateSchema(p graphql.ResolveParams) models.EvaluationCriteria {
	evaluationCriteria := models.EvaluationCriteria{}

	if name, ok := p.Args["name"]; ok {
		evaluationCriteria.Name = name.(string)
	}

	if courseID, ok := p.Args["course_id"]; ok {
		evaluationCriteria.CourseID = courseID.(string)
	}

	if maxGrade, ok := p.Args["max_grade"]; ok {
		evaluationCriteria.MaxGrade = uint(maxGrade.(int))
	}

	if passGrade, ok := p.Args["pass_grade"]; ok {
		evaluationCriteria.PassGrade = uint(passGrade.(int))
	}

	if gradeLabels, ok := p.Args["grade_labels"]; ok {
		evaluationCriteria.GradeLabels = utils.ConvertStringToJsonb(gradeLabels.(string))
	}

	evaluationCriteria.SetID(p.Args["id"].(string))

	return evaluationCriteria
}
