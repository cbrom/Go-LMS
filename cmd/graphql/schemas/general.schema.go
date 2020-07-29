package schemas

import "github.com/graphql-go/graphql"

// FetchByIDArgument general fetch by id argument
var FetchByIDArgument = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
