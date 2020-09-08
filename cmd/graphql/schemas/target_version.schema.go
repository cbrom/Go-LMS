package schemas

import (
	"errors"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// TargetVersionSchema graphql schema of target version model
var TargetVersionSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TargetVersion",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					tv := p.Source.(*models.TargetVersion)
					return tv.GetID(), nil
				},
			},
			"version_name": &graphql.Field{
				Type: graphql.String,
			},
			"content_blocks": &graphql.Field{
				Type:    graphql.NewList(ContentBlockSchema),
				Args:    FetchByIDArgument,
				Resolve: GetContentBlocks,
			},
		},
	})

// CreateTargetVersionSchema contains fields to create a new target version
var CreateTargetVersionSchema = graphql.FieldConfigArgument{
	"version_name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"target_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// TargetVersionFromSchema is an adapter for target version
func TargetVersionFromSchema(p graphql.ResolveParams) models.TargetVersion {
	targetVersion := models.TargetVersion{
		VersionName: p.Args["version_name"].(string),
		TargetID:    p.Args["target_id"].(string),
	}

	return targetVersion
}

// UpdateTargetVersionSchema contains fields to update a target version
var UpdateTargetVersionSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"version_name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"target_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// TargetVersionFromUpdateSchema is an adapter for target version
func TargetVersionFromUpdateSchema(p graphql.ResolveParams) models.TargetVersion {
	targetVersion := models.TargetVersion{}

	if versionName, ok := p.Args["version_name"]; ok {
		targetVersion.VersionName = versionName.(string)
	}
	if targetID, ok := p.Args["target_id"]; ok {
		targetVersion.TargetID = targetID.(string)
	}
	targetVersion.SetID(p.Args["id"].(string))

	return targetVersion
}

// GetContentBlocks returns a list of content blocks of a target version
func GetContentBlocks(p graphql.ResolveParams) (interface{}, error) {
	targetVersion := p.Source.(*models.TargetVersion)
	if idQuery, ok := p.Args["id"].(string); ok {
		contentBlock := models.ContentBlock{}
		contentBlock.SetID(idQuery)
		contentBlock.FetchByID()
		contentBlock.GetTargetVersion()
		if contentBlock.TargetVersion.GetID() == targetVersion.GetID() {
			return models.ContentBlockList{&contentBlock}, nil
		}
		return nil, errors.New("Content block doesn't belong to target version")
	}

	targetVersion.GetContentBlocks()
	return targetVersion.ContentBlocks, nil
}
