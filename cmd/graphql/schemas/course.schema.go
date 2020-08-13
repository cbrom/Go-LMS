package schemas

import (
	"errors"
	"fmt"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/utils"
	"time"

	"github.com/graphql-go/graphql"
)

// CourseSchema graphql schema of Course model
var CourseSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Course",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c := p.Source.(*models.Course)
					return c.GetID(), nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"ends_at": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"enable_leadboard": &graphql.Field{
				Type: graphql.Boolean,
			},
			"public_signup": &graphql.Field{
				Type: graphql.Boolean,
			},
			"featured": &graphql.Field{
				Type: graphql.Boolean,
			},
			"about": &graphql.Field{
				Type: graphql.String,
			},
			"progression_behaviour": &graphql.Field{
				Type: graphql.String,
			},
			"progression_limit": &graphql.Field{
				Type: graphql.Int,
			},
			"levels": &graphql.Field{
				Args:    FetchByIDArgument,
				Type:    graphql.NewList(LevelSchema),
				Resolve: GetCourseLevel,
			},
			"authors": &graphql.Field{
				Args:    FetchByIDArgument,
				Type:    graphql.NewList(UserSchema),
				Resolve: GetCourseAuthors,
			},
			"evaluation_criterias": &graphql.Field{
				Type: graphql.NewList(EvaluationCriteriaSchema),
			},
			"students": &graphql.Field{
				Args:    FetchByIDArgument,
				Type:    graphql.NewList(UserSchema),
				Resolve: GetStudents,
			},
			"certificates": &graphql.Field{
				Type: graphql.NewList(CertificateSchema),
			},
		},
	})

// CreateCourseSchema contains fields to create a new user
var CreateCourseSchema = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"ends_at": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"enable_leadboard": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"public_signup": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"featured": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"about": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"progression_behaviour": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"progression_limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

// CourseFromSchema course schema adapter returns course from course schema
func CourseFromSchema(p graphql.ResolveParams) models.Course {
	endsAtArg := p.Args["ends_at"]
	var endsAt *time.Time
	switch endsAtArg.(type) {
	case string:
		endsAt = utils.GetTimeFromStamp(endsAtArg.(string))
	case time.Time:
		endsAt = endsAtArg.(*time.Time)
	}
	course := models.Course{
		Name:                p.Args["name"].(string),
		EndsAt:              endsAt,
		Description:         p.Args["description"].(string),
		EnableLeadboard:     p.Args["enable_leadboard"].(bool),
		PublicSignup:        p.Args["public_signup"].(bool),
		Featured:            p.Args["featured"].(bool),
		About:               p.Args["about"].(string),
		ProgressionBehavior: p.Args["progression_behaviour"].(string),
		ProgressionLimit:    p.Args["progression_limit"].(int),
	}

	return course
}

// GetCourseLevel returns levels of a course
func GetCourseLevel(p graphql.ResolveParams) (interface{}, error) {
	course := p.Source.(*models.Course)
	if idQuery, ok := p.Args["id"].(string); ok {
		level := models.Level{}
		level.SetID(idQuery)
		level.FetchByID()
		level.GetCourse()
		if level.Course.GetID() == course.GetID() {
			// course.Levels = models.LevelList{&level}
			return models.LevelList{&level}, nil
		}
		return nil, errors.New("Level doesn't belong to course")

	}
	course.GetLevels()
	return course.Levels, nil
}

// GetCourseAuthors returns authors of a course
func GetCourseAuthors(p graphql.ResolveParams) (interface{}, error) {
	course := p.Source.(*models.Course)
	if idQuery, ok := p.Args["id"].(string); ok {
		author := models.User{}
		author.SetID(idQuery)
		author.FetchByID()

	}
	course.GetCourseAuthors()
	fmt.Printf("%v", course.Authors)
	return course.CourseAuthors, nil
}

// GetStudents returns authors of a course
func GetStudents(p graphql.ResolveParams) (interface{}, error) {
	course := p.Source.(*models.Course)
	course.GetCourseStudents()
	fmt.Printf("%v", course.AllStudents)
	return course.AllStudents, nil
}
