package resolvers

import (
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

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

// UpdateCourseAuthor updates an existing course author
func UpdateCourseAuthor (p graphql.ResolveParams) (interface{}, error) {
	courseAuthor := schemas. CourseAuthorFromUpdateSchema(p)
	if err := courseAuthor.UpdateOne(); err == nil {
		return courseAuthor.GetID(), nil
	}

	return nil, errors.New("Unable to update course author ")
}

// DeleteAuthor deletes an author from a course
func DeleteAuthor(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		author := &models.CourseAuthor{}
		author.SetID(idQuery)
		err := author.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Author id not provided")
}
