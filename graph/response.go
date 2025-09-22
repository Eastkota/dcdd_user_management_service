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
		"data":  &graphql.Field{Type: ExistUser},
		"error": &graphql.Field{Type: AuthError},
	},
})

var SingleUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var UserProfileResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserProfileResult},
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

var DcddUsersResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUsersResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{Type: DccddUserResult},
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

var SchoolResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SchoolResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: SchoolResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var GradeResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GradeResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GradeResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var EccdResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "EccdResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: EccdResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var DzongkhagResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DzongkhagResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DzongkhagResult},
		"error": &graphql.Field{Type: AuthError},
	},
})