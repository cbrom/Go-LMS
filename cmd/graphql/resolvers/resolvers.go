package resolvers

import (
	"fmt"
	"go-lms-of-pupilfirst/cmd/graphql/schemas"
	"go-lms-of-pupilfirst/cmd/models"

	"go-lms-of-pupilfirst/pkg/auth"
	"go-lms-of-pupilfirst/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var (
	authenticator *auth.Authenticator
)

// ApplyResolvers applies root queries to graphql server
func ApplyResolvers(r *gin.Engine, db *gorm.DB, auth *auth.Authenticator) {
	models.SetRepoDB(db)
	authenticator = auth

	var rootQuery = graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Query",
			Description: "User type query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type:        schemas.UserSchema,
					Description: "Returns a user by ID",
					Args:        schemas.FetchByIDArgument,
					Resolve:     GetUser,
				},
				"student": &graphql.Field{
					Type:        schemas.StudentSchema,
					Description: "Returns a student by ID",
					Args:        schemas.FetchByIDArgument,
					Resolve:     GetStudent,
				},
				"course": &graphql.Field{
					Type:        schemas.CourseSchema,
					Description: "Returns a course by ID",
					Args:        schemas.FetchByIDArgument,
					Resolve:     GetCourse,
				},
				"signin": &graphql.Field{
					Type: graphql.String,
					Args: graphql.FieldConfigArgument{
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: SignIn,
				},
			},
		})
	var rootMutation = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				/* Signup user
				 */
				"signup": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateUserSchema,
					Description: "Register new user",
					Resolve:     SignUp,
				},

				"createCourse": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateCourseSchema,
					Description: "Create a new course",
					Resolve:     CreateCourse,
				},
				"createLevel": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateLevelSchema,
					Description: "Create a new level",
					Resolve:     CreateLevel,
				},
				"enrollStudent": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateStudentCourseSchema,
					Description: "Enroll a student to a given course",
					Resolve:     EnrollStudentInCourse,
				},
				"createTargetGroup": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateTargetGroupSchema,
					Description: "Create a new target group",
					Resolve:     CreateTargetGroup,
				},
				"createAuthor": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateCourseAuthorSchema,
					Description: "Assign an author to a course",
					Resolve:     CreateAuthor,
				},
				"createCertificate": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateCertificateSchema,
					Description: "Create a new certificate",
					Resolve:     CreateCertificate,
				},
				"createContentBlock": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateContentBlockSchema,
					Description: "Create a new content block",
					Resolve:     CreateContentBlock,
				},
				"createEvaluationCriteria": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateEvaluationCriteriaSchema,
					Description: "Create a course evaluation criteria",
					Resolve:     CreateEvaluationCriteria,
				},
				"createQuiz": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateQuizSchema,
					Description: "Create a new quiz",
					Resolve:     CreateQuiz,
				},
				"createQuizQuestion": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateQuizQuestionSchema,
					Description: "Create a new quiz question",
					Resolve:     CreateQuizQuestion,
				},
				"createAnswerOption": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateAnswerOptionSchema,
					Description: "Create a new answer option",
					Resolve:     CreateAnswerOption,
				},
				"createTarget": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateTargetSchema,
					Description: "Create a new target",
					Resolve:     CreateTarget,
				},
				"createTargetVersion": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateTargetVersionSchema,
					Description: "Create a new target version",
					Resolve:     CreateTargetVersion,
				},
			},
		})

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery,
			Mutation: rootMutation,
		},
	)

	r.POST("/graphql", middlewares.JWTAuthMiddleware(authenticator), func(ctx *gin.Context) {
		var query struct {
			Query string
		}
		ctx.BindJSON(&query)
		result := executeQuery(query.Query, schema, ctx)
		ctx.JSON(200, result)
	})
}

func executeQuery(query string, schema graphql.Schema, ctx *gin.Context) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       ctx,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %+v", result.Errors)
	}
	return result
}
