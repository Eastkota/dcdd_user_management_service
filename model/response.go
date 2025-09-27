package model

type GenericAuthSuccessData struct {
    Message string `json:"message"`
    Code    string `json:"code"`
}
type DcddUserData struct {
    Users       *DcddUser              `json:"user"`
    UserProfiles *UserProfile          `json:"user_profile"`
    Error       *UserError            `json:"error"`
}

type CreateUserSuccessData struct {
    User       *DcddUser              `json:"user"`
    Profile    *UserProfile   `json:"user_profile"`
}
type DeleteUserResult struct {
    User       *DcddUser              `json:"user"`
}
type BulkSuccessResult struct {
    Message     string      `json:"message"`
    Code     string      `json:"user"`
}

type DcddGenericUserResponse struct {
	Data  interface{} `json:"data"`
	Error *UserError  `json:"error"`
}

type DcddUserProfileResult struct {
	UserProfile *UserProfile `json:"user_profile"`
}