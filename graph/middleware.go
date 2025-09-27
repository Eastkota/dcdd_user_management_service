package schema

import (
	"dcdd_user_management_service/helpers"
	"dcdd_user_management_service/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

func AuthMiddleware(next func(p graphql.ResolveParams) *model.DcddGenericUserResponse) func(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	return func(p graphql.ResolveParams) *model.DcddGenericUserResponse {
		ctx := p.Context
		user := ctx.Value(model.UserKey).(*model.DcddUser)
		if user == nil {
			return helpers.FormatError(fmt.Errorf("UnAuthorized"))
		}
		return next(p)
	}
}
