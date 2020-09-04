package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateEvaluationCriteria creates a new evaluationCriteria
func CreateEvaluationCriteria(p graphql.ResolveParams) (interface{}, error) {
	evaluationCriteria := schemas.EvaluationCriteriaFromSchema(p)
	if err := evaluationCriteria.Create(); err == nil {
		return evaluationCriteria.GetID(), nil
	}

	return nil, errors.New("Unable to create evaluationCriteria")
}

// DeleteEvaluationCriteria deletes an existing evaluation criteria
func DeleteEvaluationCriteria(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		evaluationCriteria := &models.EvaluationCriteria{}
		evaluationCriteria.SetID(idQuery)
		err := evaluationCriteria.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Evaluation criteria id not provided")
}