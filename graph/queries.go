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
						Name:    "dcddusermanagementservice",
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
			"FetchProfileByDcddUserId": &graphql.Field{
				Type: DcddUserProfileResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchProfileByDcddUserId(p), nil
				},
			},
			"fetchAllDcddUsers": &graphql.Field{
				Type: DcddUsersByDateRangeResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchAllUsers(p)
				},
			},
			"fetchAllActiveDcddUsers": &graphql.Field{
				Type: DcddUsersByDateRangeResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchAllActiveUsers(p)
				},
			},
			"FetchDcddUsersByDateRange": &graphql.Field{
				Type: DcddUsersByDateRangeResponse,
				Args: graphql.FieldConfigArgument{
					"fromDate": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"toDate":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: resolver.FetchDcddUsersByDateRange,
			},


			"fetchSchools": &graphql.Field{
				Type: SchoolResponse,
				Args: graphql.FieldConfigArgument{
					"dzongkhag_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchSchool(p), nil
				},
			},
			"fetchEccd": &graphql.Field{
				Type: EccdResponse,
				Args: graphql.FieldConfigArgument{
					"dzongkhag_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchEccd(p), nil
				},
			},
			"fetchDzongkhag": &graphql.Field{
				Type: DzongkhagResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchDzongkhag(p), nil
				},
			},
			"fetchGrade": &graphql.Field{
				Type: GradeResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchGrade(p), nil
				},
			},
		},
	})
}
