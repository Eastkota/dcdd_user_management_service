package schema

import "github.com/graphql-go/graphql"


var GenericAuthResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAuthResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: AuthGenericSuccessData},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var DcddUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{Type: graphql.NewList(DcddUser)},
		"error": &graphql.Field{Type: DcddUserError},
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
		"data":  &graphql.Field{Type: DcddUserResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var DcddUsersByDateRangeResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUsersByDateRangeResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{Type: graphql.NewList(DcddUserResult)},
		"error": &graphql.Field{Type: DcddAuthError },
	},
})


var DcddUserStatusResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserStatusResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DcddUserStatus},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var SchoolResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SchoolResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: SchoolResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var GradeResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GradeResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GradeResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var EccdResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "EccdResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: EccdResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})

var DzongkhagResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "DzongkhagResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: DzongkhagResult},
		"error": &graphql.Field{Type: DcddAuthError},
	},
})