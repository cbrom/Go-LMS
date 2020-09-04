package resolvers

import (
	"fmt"
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

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

// UpdateTarget updates an existing target
func UpdateTarget(p graphql.ResolveParams) (interface{}, error) {
	target := schemas.TargetFromUpdateSchema(p)
	if err := target.UpdateOne(); err == nil {
		return target.GetID(), nil
	}

	return nil, errors.New("Unable to update target ")
}

// DeleteTarget deletes an existing target
func DeleteTarget(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string) 
	if ok {
		target := &models.Target{}
		target.SetID(idQuery)
		err := target.SoftDelete()
		return nil, err
	}

	return nil, errors.New("Target id not provided")
}

// CreateTargetVersion creates a new targetVersion
func CreateTargetVersion(p graphql.ResolveParams) (interface{}, error) {
	targetVersion := schemas.TargetVersionFromSchema(p)
	if err := targetVersion.Create(); err == nil {
		return targetVersion.GetID(), nil
	}

	return nil, errors.New("Unable to create targetVersion")
}

// UpdateTargetVersion updates an existing target version
func UpdateTargetVersion(p graphql.ResolveParams) (interface{}, error) {
	targetVersion := schemas.TargetVersionFromUpdateSchema(p)
	if err := targetVersion.UpdateOne(); err == nil {
		return targetVersion.GetID(), nil
	}

	return nil, errors.New("Unable to update target version ")
}

// DeleteTargetVersion deletes an existing target version
func DeleteTargetVersion(p graphql.ResolveParams) (interface{}, error) {
	idQuery, ok := p.Args["id"].(string)
	if ok {
		targetVersion := &models.TargetVersion{}
		targetVersion.SetID(idQuery)
		err := targetVersion.SoftDelete() 
		return nil, err
	}

	return nil, errors.New("Target version id not provided")
}