package schema

import (
	"dcdd_user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// Define UserType
var DcddUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUser",
	Fields: graphql.Fields{
		"id":              &graphql.Field{Type: scalar.UUID},
		"user_identifier": &graphql.Field{Type: graphql.String},
		"email":           &graphql.Field{Type: graphql.String},
		"mobile_no":       &graphql.Field{Type: graphql.String},
		"login_id":        &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
		"password":        &graphql.Field{Type: graphql.String},
		"student_id":      &graphql.Field{Type: graphql.String},
		"category":        &graphql.Field{Type: graphql.String},
		"created_at":      &graphql.Field{Type: scalar.Time},
		"updated_at":      &graphql.Field{Type: scalar.Time},
	},
})

// Define UserType
var DcddUserProfile = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserProfile",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"student_id":                 &graphql.Field{Type: scalar.UUID},
		"grade_id":                   &graphql.Field{Type: scalar.UUID},
		"eccd_id":                    &graphql.Field{Type: scalar.UUID},
		"dzongkhag_idid":             &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"gender":                     &graphql.Field{Type: graphql.String},
		"cid":                     &graphql.Field{Type: graphql.String},
		"dob":                     &graphql.Field{Type: scalar.Time},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})

var UserProfileAndUsers = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileAndUsers",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"student_id":                 &graphql.Field{Type: scalar.UUID},
		"user_identifier":			  &graphql.Field{Type: graphql.String},
		"email":           			  &graphql.Field{Type: graphql.String},
		"mobile_no":       			  &graphql.Field{Type: graphql.String},
		"login_id":        			  &graphql.Field{Type: graphql.String},
		"status":            		  &graphql.Field{Type: graphql.String},
		"password":        			  &graphql.Field{Type: graphql.String},
		"category":        		      &graphql.Field{Type: graphql.String},
		"grade_id":                   &graphql.Field{Type: scalar.UUID},
		"eccd_id":                    &graphql.Field{Type: scalar.UUID},
		"dzongkhag_idid":             &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"gender":                     &graphql.Field{Type: graphql.String},
		"cid":                     	  &graphql.Field{Type: graphql.String},
		"dob":                     	  &graphql.Field{Type: scalar.Time},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})


var DcddExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"user_id":    &graphql.Field{Type: scalar.UUID},
	},
})
// var DcddUserResult = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "DcddUserResult",
// 	Fields: graphql.Fields{
// 		"user": &graphql.Field{Type: DcddUser},
// 	},
// })

var DccddUserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "DccddUserResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: User},
		"user_profile": &graphql.Field{Type: UserProfile},
	},
})

var UserProfileResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResult",
	Fields: graphql.Fields{
		"user_profile": &graphql.Field{Type: DcddUserProfile},
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

var DcddUserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserResult",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: DcddUser},
		"profile":    &graphql.Field{Type: DcddUserProfile},
	},
})
var DcddUserStatus = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserStatus",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: DcddUser},
	},
})

var AuthSchool = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthSchool",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: scalar.UUID},
		"name":       &graphql.Field{Type: graphql.String},
		"message":    &graphql.Field{Type: graphql.String},
		"pvt_public": &graphql.Field{Type: graphql.String},
	},
})
var AuthGrade = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthGrade",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: scalar.UUID},
		"name": &graphql.Field{Type: graphql.String},
	},
})
var AuthEccd = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthEccd",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: scalar.UUID},
		"name":       &graphql.Field{Type: graphql.String},
		"sort":       &graphql.Field{Type: graphql.String},
	},
})
var AuthDzongkhag = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthDzongkhag",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: scalar.UUID},
		"name": &graphql.Field{Type: graphql.String},
	},
})

