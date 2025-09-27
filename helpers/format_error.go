package helpers

import "dcdd_user_management_service/model"

func FormatError(err error) *model.DcddGenericUserResponse {
	return &model.DcddGenericUserResponse{
		Data: nil,
		Error: &model.UserError{
			Message: err.Error(),
		},
	}
}
