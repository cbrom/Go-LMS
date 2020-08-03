package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

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
