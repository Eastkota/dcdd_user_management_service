package schema

import (
	"dcdd_user_management_service/helpers"
	"dcdd_user_management_service/model"
	"dcdd_user_management_service/resolver"
	"dcdd_user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

func NewQueryType(resolver *resolvers.UserResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"service": &graphql.Field{
				Type: graphql.NewNonNull(Service),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					schema, err := GetSchema()
					if err != nil {
						return nil, err
					}

					serviceInfo := model.Service{
						Name:    "AuthService",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
			"CheckForDcddExistingUser": &graphql.Field{
				Type: CheckForDcddExistingUserResponse,
				Args: graphql.FieldConfigArgument{
					"field": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"value": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.CheckForDcddExistingUser(p), nil
				},
			},
			"fetchProfileByUserId": &graphql.Field{
				Type: UserProfileResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchProfileByUserId(p), nil
				},
			},
			"fetchAllUsers": &graphql.Field{
				Type: graphql.NewList(UserProfileAndUsers),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchAllUsers(p)
				},
			},
			"fetchAllActiveUsers": &graphql.Field{
				Type: graphql.NewList(UserProfileAndUsers),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchAllActiveUsers(p)
				},
			},
			"fetchUsersByDateRange": &graphql.Field{
				Type: graphql.NewList(UserProfileAndUsers), // or just DcddUser type if you want only users
				Args: graphql.FieldConfigArgument{
					"fromDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String), // YYYY-MM-DD
					},
					"toDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchUsersByDateRange(p)
				},
			},
		},
	})
}
