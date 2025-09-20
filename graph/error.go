package schema

import "github.com/graphql-go/graphql"

var DcddAuthError = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddAuthError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})

var DcddUserError = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
