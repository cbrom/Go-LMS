package schemas

import (
	"errors"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/utils"
	"time"

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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					t := p.Source.(*models.Target)
					checkList := t.CheckList
					return string(checkList.RawMessage), nil
				},
			},
			"review_checklist": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					t := p.Source.(*models.Target)
					reviewChecklist := t.ReviewChecklist
					return string(reviewChecklist.RawMessage), nil
				},
			},
			"target_versions": &graphql.Field{
				Type:    graphql.NewList(TargetVersionSchema),
				Args:    FetchByIDArgument,
				Resolve: GetTargetVersions,
			},
			"quizzes": &graphql.Field{
				Type:    graphql.NewList(QuizSchema),
				Args:    FetchByIDArgument,
				Resolve: GetQuizzes,
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

// TargetFromSchema is an adapter for target
func TargetFromSchema(p graphql.ResolveParams) models.Target {
	sessionAtArg := p.Args["session_at"]
	var sessionAt *time.Time
	switch sessionAtArg.(type) {
	case string:
		sessionAt = utils.GetTimeFromStamp(sessionAtArg.(string))
	case time.Time:
		sessionAt = sessionAtArg.(*time.Time)
	}
	checkList := utils.ConvertStringToJsonb(p.Args["check_list"].(string))
	checkListReview := utils.ConvertStringToJsonb(p.Args["review_checklist"].(string))

	target := models.Target{
		TargetGroupID:          p.Args["target_group_id"].(string),
		Role:                   p.Args["role"].(string),
		Title:                  p.Args["title"].(string),
		Description:            p.Args["description"].(string),
		CompletionInstructions: p.Args["completion_instructions"].(string),
		ResourceURL:            p.Args["resource_url"].(string),
		SortIndex:              p.Args["sort_index"].(int),
		SessionAt:              sessionAt,
		LinkToComplete:         p.Args["link_to_complete"].(string),
		Resubmittable:          p.Args["resubmittable"].(bool),
		CheckList:              checkList,
		ReviewChecklist:        checkListReview,
	}

	return target
}

// UpdateTargetSchema contains fields to update a target
var UpdateTargetSchema = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
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

// TargetFromUpdateSchema is an adapter for target
func TargetFromUpdateSchema(p graphql.ResolveParams) models.Target {

	target := models.Target{}

	if targetGroupID, ok := p.Args["target_group_id"]; ok {
		target.TargetGroupID = targetGroupID.(string)
	}
	if role, ok := p.Args["role"]; ok {
		target.Role = role.(string)
	}
	if title, ok := p.Args["title"]; ok {
		target.Title = title.(string)
	}
	if description, ok := p.Args["description"]; ok {
		target.Description = description.(string)
	}
	if completionInstructions, ok := p.Args["completion_instructions"]; ok {
		target.CompletionInstructions = completionInstructions.(string)
	}
	if resourceURL, ok := p.Args["resource_url"]; ok {
		target.ResourceURL = resourceURL.(string)
	}
	if sortIndex, ok := p.Args["sort_index"]; ok {
		target.SortIndex = sortIndex.(int)
	}
	if sessionAtArg, ok := p.Args["session_at"]; ok {
		var sessionAt *time.Time
		switch sessionAtArg.(type) {
		case string:
			sessionAt = utils.GetTimeFromStamp(sessionAtArg.(string))
		case time.Time:
			sessionAt = sessionAtArg.(*time.Time)
		}
		target.SessionAt = sessionAt
	}
	if linkToComplete, ok := p.Args["link_to_complete"]; ok {
		target.LinkToComplete = linkToComplete.(string)
	}
	if resubmittable, ok := p.Args["resubmittable"]; ok {
		target.Resubmittable = resubmittable.(bool)
	}
	if checkListArg, ok := p.Args["check_list"]; ok {
		checkList := utils.ConvertStringToJsonb(checkListArg.(string))
		target.CheckList = checkList
	}
	if reviewChecklistArg, ok := p.Args["review_checklist"]; ok {
		checkListReview := utils.ConvertStringToJsonb(reviewChecklistArg.(string))
		target.ReviewChecklist = checkListReview
	}
	target.SetID(p.Args["id"].(string))

	return target
}

// GetTargetVersions returns a list of target versions of the target
func GetTargetVersions(p graphql.ResolveParams) (interface{}, error) {
	target := p.Source.(*models.Target)
	if idQuery, ok := p.Args["id"].(string); ok {
		targetVersion := models.TargetVersion{}
		targetVersion.SetID(idQuery)
		targetVersion.FetchByID()
		targetVersion.GetTarget()
		if targetVersion.Target.GetID() == target.GetID() {
			return models.TargetVersionList{&targetVersion}, nil
		}

		return nil, errors.New("target version doesn't velong to target")
	}

	target.GetVersions()
	return target.TargetVersions, nil
}

// GetQuizzes returns a list of quizzes of the target
func GetQuizzes(p graphql.ResolveParams) (interface{}, error) {
	target := p.Source.(*models.Target)
	if idQuery, ok := p.Args["id"].(string); ok {
		quiz := models.Quiz{}
		quiz.SetID(idQuery)
		quiz.FetchByID()
		quiz.GetTarget()
		if quiz.Target.GetID() == target.GetID() {
			return models.QuizList{&quiz}, nil
		}

		return nil, errors.New("quiz doesn't velong to target")
	}

	target.GetQuizzes()
	return target.Quizzes, nil
}
