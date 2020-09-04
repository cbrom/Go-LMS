package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

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

// DeleteContentBlock deletes an existing content block
func DeleteContentBlock(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		contentBlock := &models.ContentBlock{}
		contentBlock.SetID(idQuery)
		err := contentBlock.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Content block id not provided")
}
