package resolvers

import (
	"fmt"
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateTarget creates a new target
func CreateTarget(p graphql.ResolveParams) (interface{}, error) {
	target := schemas.TargetFromSchema(p)
	fmt.Printf("target to be created %+v \n\n", target)
	if err := target.Create(); err == nil {
		return target.GetID(), nil
	}

	return nil, errors.New("Unable to create target")
}

// CreateTargetVersion creates a new targetVersion
func CreateTargetVersion(p graphql.ResolveParams) (interface{}, error) {
	targetVersion := schemas.TargetVersionFromSchema(p)
	if err := targetVersion.Create(); err == nil {
		return targetVersion.GetID(), nil
	}

	return nil, errors.New("Unable to create targetVersion")
}
