package schema

import (
	"dcdd_user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

var SignupInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "DcddSignupInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"mobile_no": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"gender": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"student_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"grade_id": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"school_id": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"eccd_id": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"cid": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"dob": &graphql.InputObjectFieldConfig{
				Type: scalar.Time,
			},
			"dzongkhag_id": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var UpdatePasswordInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdatePasswordInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"user_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"confirm_password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"current_password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var UpdateSingleAuthDataInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdateSingleAuthDataInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"field": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"value": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var ResetPasswordInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ResetPasswordInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"user_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(scalar.UUID),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"confirm_password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var UserProfileInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":            &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"profile_picture": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"gender":          &graphql.InputObjectFieldConfig{Type: graphql.String},
		"dzongkhag_id":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(scalar.UUID),},
		"user_id":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(scalar.UUID)},
	},
})
