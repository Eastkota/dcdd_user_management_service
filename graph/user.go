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
		"dzongkhag_id":               &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"gender":                     &graphql.Field{Type: graphql.String},
		"cid":                     	  &graphql.Field{Type: graphql.String},
		"dob":                        &graphql.Field{Type: scalar.Time},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})

var UserProfileAndUsers = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileAndUsers",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type:DcddUser},
		"user_profile": &graphql.Field{Type: DcddUserProfile},
	},
})


var DcddExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"user_id":    &graphql.Field{Type: scalar.UUID},
	},
})

var DcddUserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: DcddUser},
		"user_profile": &graphql.Field{Type: DcddUserProfile},
	},
})

var DcddUserProfileResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "DcddUserProfileResult",
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

var GenericDcddUserSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericDcddUserSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
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
var School = graphql.NewObject(graphql.ObjectConfig{
	Name: "School",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.String},
		"name":       &graphql.Field{Type: graphql.String},
		"pvt_public":    &graphql.Field{Type: graphql.String},
	},
})
var SchoolResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "SchoolResult",
	Fields: graphql.Fields{
		"school": &graphql.Field{Type: graphql.NewList(School)},
	},
})
var Grade = graphql.NewObject(graphql.ObjectConfig{
	Name: "Grade",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: scalar.UUID},
		"name": &graphql.Field{Type: graphql.String},
	},
})
var GradeResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "GradeResult",
	Fields: graphql.Fields{
		"grade": &graphql.Field{Type: graphql.NewList(Grade)},
	},
})
var Eccd = graphql.NewObject(graphql.ObjectConfig{
	Name: "Eccd",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: scalar.UUID},
		"name":       &graphql.Field{Type: graphql.String},
		"sort":       &graphql.Field{Type: graphql.String},
	},
})
var EccdResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "EccdResult",
	Fields: graphql.Fields{
		"eccd": &graphql.Field{Type: graphql.NewList(Eccd)},
	},
})

var Dzongkhag = graphql.NewObject(graphql.ObjectConfig{
	Name: "Dzongkhag",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: scalar.UUID},
		"name": &graphql.Field{Type: graphql.String},
	},
})
var DzongkhagResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "DzongkhagResult",
	Fields: graphql.Fields{
		"dzongkhags": &graphql.Field{Type: graphql.NewList(Dzongkhag)},
	},
})
