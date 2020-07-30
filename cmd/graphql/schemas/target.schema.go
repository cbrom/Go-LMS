package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/graphql-go/graphql"
)

// TargetSchema graphql schema of target  model
var TargetSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Target",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					t := p.Source.(*models.Target)
					return t.GetID(), nil
				},
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"completion_instructions": &graphql.Field{
				Type: graphql.String,
			},
			"resource_url": &graphql.Field{
				Type: graphql.String,
			},
			"sort_index": &graphql.Field{
				Type: graphql.Int,
			},
			"session_at": &graphql.Field{
				Type: graphql.String,
			},
			"link_to_complete": &graphql.Field{
				Type: graphql.String,
			},
			"resubmittable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"check_list": &graphql.Field{
				Type: graphql.String,
			},
			"review_checklist": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// CreateTargetSchema contains fields to create a new target
var CreateTargetSchema = graphql.FieldConfigArgument{
	"target_group_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"role": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"title": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"completion_instructions": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"resource_url": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"sort_index": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"session_at": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"link_to_complete": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"resubmittable": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"check_list": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"review_checklist": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
