package schema

import "github.com/graphql-go/graphql"


var GenericAuthResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAuthResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: AuthGenericSuccessData},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var CheckForDcddExistingUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckForDcddExistingUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddExistUser},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var SingleUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var DcddUserProfileResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserProfileResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserProfileResult},
		"error": &graphql.Field{Type: DcddUserError},
	},
})

var DcddGenericUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddGenericUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GenericDcddUserSuccessData},
		"error": &graphql.Field{Type: DcddUserError},
	},
})

var DcddCreateUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddCreateUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUser},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var DcddUsersResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUsersResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{Type: DccddUserResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})
var DcddUserStatusResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserStatusResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserStatus},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

