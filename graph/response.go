package schema

import "github.com/graphql-go/graphql"


var GenericAuthResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAuthResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: AuthGenericSuccessData},
		"error": &graphql.Field{Type: AuthError},
	},
})

var CheckForDcddExistingUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckForDcddExistingUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddExistUser},
		"error": &graphql.Field{Type: AuthError},
	},
})

var SingleUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var UserProfileResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserProfileResult},
		"error": &graphql.Field{Type: UserError},
	},
})

var GenericUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GenericUserSuccessData},
		"error": &graphql.Field{Type: UserError},
	},
})

var CreateUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUser},
		"error": &graphql.Field{Type: AuthError},
	},
})
var UserStatusResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserStatusResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserStatus},
		"error": &graphql.Field{Type: AuthError},
	},
})

