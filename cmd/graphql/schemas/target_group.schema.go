package schemas

import (
	"errors"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// TargetGroupSchema graphql schema of TargetGroup model
var TargetGroupSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TargetGroup",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					tg := p.Source.(*models.TargetGroup)
					return tg.GetID(), nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"sort_index": &graphql.Field{
				Type: graphql.Int,
			},
			"milestone": &graphql.Field{
				Type: graphql.Boolean,
			},
			"targets": &graphql.Field{
				Type:    graphql.NewList(TargetSchema),
				Args:    FetchByIDArgument,
				Resolve: GetTargets,
			},
		},
	})

// CreateTargetGroupSchema contains fields used to create a new target group
var CreateTargetGroupSchema = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"sort_index": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"milestone": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"level_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// TargetGroupFromSchema is an adapter for target group
func TargetGroupFromSchema(p graphql.ResolveParams) models.TargetGroup {
	targetGroup := models.TargetGroup{
		Name:        p.Args["name"].(string),
		Description: p.Args["description"].(string),
		SortIndex:   p.Args["sort_index"].(int),
		Milestone:   p.Args["milestone"].(bool),
		LevelID:     p.Args["level_id"].(string),
	}

	return targetGroup
}

// GetTargets returns a list of targets of a target group
func GetTargets(p graphql.ResolveParams) (interface{}, error) {
	targetGroup := p.Source.(*models.TargetGroup)
	if idQuery, ok := p.Args["id"].(string); ok {
		target := models.Target{}
		target.SetID(idQuery)
		target.FetchByID()
		target.GetTargetGroup()
		if target.TargetGroup.GetID() == targetGroup.GetID() {
			return models.TargetList{&target}, nil
		}
		return nil, errors.New("target doesn't belong to target group")
	}

	targetGroup.GetTargets()
	return targetGroup.Targets, nil
}
