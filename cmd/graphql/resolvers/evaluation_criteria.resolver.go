package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

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
