package model

type GenericAuthSuccessData struct {
    Message string `json:"message"`
    Code    string `json:"code"`
}
type DcddUserData struct {
    User       *DcddUser              `json:"user"`
    UserProfile *UserProfile          `json:"user_profile"`
}

type CreateUserSuccessData struct {
    User       *DcddUser              `json:"user"`
    Profile    *UserProfile   `json:"profile"`
}
type DeleteUserResult struct {
    User       *DcddUser              `json:"user"`
}
type BulkSuccessResult struct {
    Message     string      `json:"message"`
    Code     string      `json:"user"`
}

type GenericUserResponse struct {
	Data  interface{}
	Error *UserError
}

type DcddUserProfileResult struct {
	UserProfile *UserProfile `json:"user_profile"`
}