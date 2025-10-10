package helpers

import (
	"dcdd_user_management_service/config"
	"dcdd_user_management_service/model"

	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func DcddValidateToken(tokenStr string) (*model.DcddUser, error) {
	authServiceClient := graphql.NewClient(config.AuthServiceApi())
	req := graphql.NewRequest(`
		query DcddValidateToken($input:  String){
			dcddValidateToken(token: $input) {
				data {
					user {
						category
						created_at
						email
						id
						login_id
						mobile_no
						password
						status
						student_id
						updated_at
						user_identifier
					}
				}
				error {
					code
					field
					message
				}
			}
		}
	`)
	req.Var("input", tokenStr)
	req.Header.Set("Cache-Control", "no-cache")

	var response struct {
		DcddValidateToken struct {
			Data struct {
				User model.DcddUser `json:"user"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"dcddValidateToken"`
	}

	err := authServiceClient.Run(context.Background(), req, &response)
	if err != nil {
		return nil, fmt.Errorf("Invalid_Token")
	}
	if response.DcddValidateToken.Error.Message != "" {
		return nil, fmt.Errorf(response.DcddValidateToken.Error.Message)
	}
	return &response.DcddValidateToken.Data.User, err
}