package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateTargetGroup creates a new target group
func CreateTargetGroup(p graphql.ResolveParams) (interface{}, error) {
	targetGroup := schemas.TargetGroupFromSchema(p)
	if err := targetGroup.Create(); err == nil {
		return targetGroup.GetID(), nil
	}

	return nil, errors.New("Unable to create target group")
}

// DeleteTargetGroup deletes an existing target group
func DeleteTargetGroup(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		targetGroup := &models.TargetGroup{}
		targetGroup.SetID(idQuery)
		err := targetGroup.SoftDelete()
		return nil, err
	}
	return nil, errors.New("TargetGroup id not provided")
}