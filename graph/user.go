package schema

import (
	"dcdd_user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// Define UserType
var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":              &graphql.Field{Type: scalar.UUID},
		"user_identifier": &graphql.Field{Type: graphql.String},
		"name":            &graphql.Field{Type: graphql.String},
		"email":           &graphql.Field{Type: graphql.String},
		"mobile_no":       &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
		"created_at":      &graphql.Field{Type: scalar.Time},
		"updated_at":      &graphql.Field{Type: scalar.Time},
	},
})

// Define UserType
var UserProfile = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfile",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"student_id":                         &graphql.Field{Type: scalar.UUID},
		"grade_id":                         &graphql.Field{Type: scalar.UUID},
		"eccd_id":                         &graphql.Field{Type: scalar.UUID},
		"dzongkhag_idid":                         &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"gender":                     &graphql.Field{Type: graphql.String},
		"cid":                     &graphql.Field{Type: graphql.String},
		"dob":                     &graphql.Field{Type: scalar.Time},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})


var ExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"user_id":    &graphql.Field{Type: scalar.UUID},
	},
})
var UserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: User},
	},
})

var UserProfileResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResult",
	Fields: graphql.Fields{
		"user_profile": &graphql.Field{Type: UserProfile},
	},
})

var AuthGenericSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthGenericSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

var GenericUserSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericUserSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

var DcddUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUser",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: User},
		"profile":    &graphql.Field{Type: UserProfile},
	},
})
var DcddUserStatus = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserStatus",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: User},
	},
})
