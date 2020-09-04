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

				/*
				Create a new course
				*/
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
				"issueCertificate": &graphql.Field{
					Type:        graphql.String,
					Args:        schemas.CreateIssuedCertificateSchema,
					Description: "Issue a certificate to a student",
					Resolve:     CreateIssuedCertificate,
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

				/*
				* Delete records
				*/

				"deleteUser": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing user",
					Resolve: DeleteUser,
				},
				"deleteCourse": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing course",
					Resolve: DeleteCourse,
				},
				"deleteLevel": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing level",
					Resolve: DeleteLevel,
				},
				"unEnrollStudent": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an enrolled student",
					Resolve: UnenrollStudent,
				},
				"deleteTargetGroup": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing target group",
					Resolve: DeleteTargetGroup,
				},
				"deleteAuthor": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing author",
					Resolve: DeleteAuthor,
				},
				"deleteCertificate": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing certificate",
					Resolve: DeleteCertificate,
				},
				"deleteIssuedCertificate": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an issued certificate",
					Resolve: UnissueCertificate,
				},
				"deleteContentBlock": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing content block",
					Resolve: DeleteContentBlock,
				},
				"deleteEvaluationCriteria": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing evaluation criteria",
					Resolve: DeleteEvaluationCriteria,
				},
				"deleteQuiz": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing quiz",
					Resolve: DeleteQuiz,
				},
				"deleteQuizQuestion": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing quiz question",
					Resolve: DeleteQuizQuestion,
				},
				"deleteAnswerOption": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing answer option",
					Resolve: DeleteAnswerOption,
				},
				"deleteTarget": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing target",
					Resolve: DeleteTarget,
				},
				"deleteTargetVersion": &graphql.Field{
					Type: graphql.String,
					Args: StringArgs,
					Description: "Delete an existing target version",
					Resolve: DeleteTargetVersion,
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

var StringArgs = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}