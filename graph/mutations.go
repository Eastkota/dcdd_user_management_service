package schema

import (
	"dcdd_user_management_service/resolver"
	"dcdd_user_management_service/graph/scalar"
	
	"github.com/graphql-go/graphql"
)

func NewMutationType(resolver *resolvers.UserResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createDcddUser": &graphql.Field{
				Type: DcddCreateUserResponse,
				Args: graphql.FieldConfigArgument{
					"signup_input": &graphql.ArgumentConfig{
						Type: SignupInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.CreateDcddUser(p), nil
				},
			},
			// "CreateDcddUserProfile": &graphql.Field{
			// 	Type: UserProfileResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"input": &graphql.ArgumentConfig{
			// 			Type: UserProfileInput,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return resolver.CreateDcddUserProfile(p), nil
			// 	},
			// },
			"updateDcddUser": &graphql.Field{
				Type: DcddCreateUserResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
					"signup_input": &graphql.ArgumentConfig{
						Type: SignupInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.UpdateDcddUser(p), nil
				},
			},
			"UpdateDcddUserStatus" : &graphql.Field{
				Type: DcddUserStatusResponse,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.UpdateDcddUserStatus(p), nil
				},
			},
			"bulkRegistration": &graphql.Field{
				Type: DcddGenericUserResponse,
				Args: graphql.FieldConfigArgument{
					"csv_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.BulkRegistration(p), nil
				},
			},
		},
	})
}
