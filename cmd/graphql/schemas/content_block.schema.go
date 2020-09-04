package schemas

import (
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/graphql-go/graphql"
)

// ContentBlockSchema graphql schema of content block model
var ContentBlockSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentBlock",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cb := p.Source.(*models.ContentBlock)
					return cb.GetID(), nil
				},
			},
			"block_type": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cb := p.Source.(*models.ContentBlock)
					content := cb.Content
					return string(content.RawMessage), nil
				},
			},
			"sort_index": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

// CreateContentBlockSchema contains fields to create a new content block
var CreateContentBlockSchema = graphql.FieldConfigArgument{
	"block_type": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"sort_index": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"target_version_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// ContentBlockFromSchema is an adapter for content block
func ContentBlockFromSchema(p graphql.ResolveParams) models.ContentBlock {
	content := utils.ConvertStringToJsonb(p.Args["content"].(string))
	contentBlock := models.ContentBlock{
		BlockType:       p.Args["block_type"].(string),
		Content:         content,
		SortIndex:       p.Args["sort_index"].(int),
		TargetVersionID: p.Args["target_version_id"].(string),
	}

	return contentBlock
}

// UpdateContentBlockSchema contains fields to update a content block
var UpdateContentBlockSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"block_type": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"sort_index": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"target_version_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// ContentBlockFromUpdateSchema is an adapter for content block
func ContentBlockFromUpdateSchema(p graphql.ResolveParams) models.ContentBlock {
	content := utils.ConvertStringToJsonb(p.Args["content"].(string))
	contentBlock := models.ContentBlock{
		BlockType:       p.Args["block_type"].(string),
		Content:         content,
		SortIndex:       p.Args["sort_index"].(int),
		TargetVersionID: p.Args["target_version_id"].(string),
	}

	contentBlock.SetID(p.Args["id"].(string))

	return contentBlock
}
