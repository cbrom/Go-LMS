package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateContentBlock creates a new contentBlock
func CreateContentBlock(p graphql.ResolveParams) (interface{}, error) {
	contentBlock := schemas.ContentBlockFromSchema(p)
	if err := contentBlock.Create(); err == nil {
		return contentBlock.GetID(), nil
	}

	return nil, errors.New("Unable to create contentBlock")
}
