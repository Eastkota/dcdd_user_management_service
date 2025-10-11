package schema

import (
	"dcdd_user_management_service/helpers"
	"dcdd_user_management_service/model"
	
	"context"
	"net/http"
	"fmt"


	"github.com/graphql-go/graphql"
)
func AuthMiddleware(next func(p graphql.ResolveParams) *model.DcddGenericUserResponse) func(p graphql.ResolveParams) *model.DcddGenericUserResponse {
    return func(p graphql.ResolveParams) *model.DcddGenericUserResponse {
        ctx := p.Context
        userInterface := ctx.Value("user")

        var user *model.DcddUser
        if userInterface == nil {
            if req, ok := ctx.Value("http_request").(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := helpers.ValidateToken(authHeader)
                if err != nil {
                    return helpers.FormatError(err)
                }
                if u == nil {
                    return helpers.FormatError(fmt.Errorf("invalid_token"))
                }

                ctx = context.WithValue(ctx, "user", u)
                p.Context = ctx
                user = u
            } else {
                return helpers.FormatError(fmt.Errorf("invalid_token"))
            }
        } else {
            user, _ = userInterface.(*model.DcddUser)
        }

        if user == nil {
            return helpers.FormatError(fmt.Errorf("invalid_token"))
        }
        return next(p)
    }
}
