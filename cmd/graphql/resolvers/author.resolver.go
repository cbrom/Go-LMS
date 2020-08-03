package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// CreateAuthor creates a new author
func CreateAuthor(p graphql.ResolveParams) (interface{}, error) {
	author := schemas.CourseAuthorFromSchema(p)
	if err := author.Create(); err == nil {
		return author.GetID(), nil
	}

	return nil, errors.New("Unable to create author")
}
