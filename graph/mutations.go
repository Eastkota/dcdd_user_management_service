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
					return AuthMiddleware(PermissionMiddleware("create",resolver.CreateDcddUser))(p), nil
				},
			},
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
					return AuthMiddleware(PermissionMiddleware("update",resolver.UpdateDcddUser))(p), nil
				},
			},
			"updateDcddUserPassword": &graphql.Field{
				Type: DcddSingleUserResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
					"DcddResetPasswordInput": &graphql.ArgumentConfig{
						Type: DcddResetPasswordInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("update",resolver.UpdateDcddUserPassword))(p), nil
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
					return AuthMiddleware(PermissionMiddleware("update",resolver.UpdateDcddUserStatus))(p), nil
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
					return AuthMiddleware(PermissionMiddleware("create",resolver.BulkRegistration))(p), nil
				},
			},
		},
	})
}
