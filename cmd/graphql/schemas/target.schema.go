package schemas

import (
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
