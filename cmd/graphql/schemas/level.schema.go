package schemas

import (
	"errors"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/utils"
	"time"

	"github.com/graphql-go/graphql"
)

// LevelSchema graphql schema of Level model
var LevelSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Level",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					l := p.Source.(*models.Level)
					return l.GetID(), nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"number": &graphql.Field{
				Type: graphql.Int,
			},
			"unlock_on": &graphql.Field{
				Type: graphql.String,
			},
			"target_groups": &graphql.Field{
				Type:    graphql.NewList(TargetGroupSchema),
				Args:    FetchByIDArgument,
				Resolve: GetLevelTargeGroup,
			},
		},
	})

// CreateLevelSchema contains fields to create a new level
var CreateLevelSchema = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"number": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"unlock_on": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"course_id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// CreateLevelFromSchema adapter for level schema and level model
func CreateLevelFromSchema(p graphql.ResolveParams) models.Level {
	unlockOnArg := p.Args["unlock_on"]
	var unlockOn *time.Time
	switch unlockOnArg.(type) {
	case string:
		unlockOn = utils.GetTimeFromStamp(unlockOnArg.(string))
	case time.Time:
		unlockOn = unlockOnArg.(*time.Time)
	}
	level := models.Level{
		Name:        p.Args["name"].(string),
		CourseID:    p.Args["course_id"].(string),
		Description: p.Args["description"].(string),
		Number:      p.Args["number"].(int),
		UnlockOn:    unlockOn,
	}

	return level
}

// GetLevelTargeGroup returns target groups of a level
func GetLevelTargeGroup(p graphql.ResolveParams) (interface{}, error) {
	level := p.Source.(*models.Level)
	if idQuery, ok := p.Args["id"].(string); ok {
		targetGroup := models.TargetGroup{}
		targetGroup.SetID(idQuery)
		targetGroup.FetchByID()
		targetGroup.GetLevel()
		if targetGroup.Level.GetID() == level.GetID() {
			return models.TargetGroupList{&targetGroup}, nil
		}
		return nil, errors.New("Target group doesn't belong to level")
	}

	level.GetTargetGroups()
	return level.TargetGroups, nil
}
