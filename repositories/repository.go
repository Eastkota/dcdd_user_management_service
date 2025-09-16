package repositories

import (
	"dcdd_user_management_service/model"

	"context"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	CheckForDcddExistingUser(field, value string) (*model.DcddUser, error)
	CreateDcddUser(signUpInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	CreateUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error)
	UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	GetAllUsers() ([]model.DcddUser, error)
	GetAllActiveUsers() ([]model.DcddUser, error)
	UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error)
	BulkRegistration(signupInputs []model.SignupInput) (error) 

	FetchProfileByUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error)
	FetchUserByLoginID(field, value string) (*model.DcddUser, error)
}
