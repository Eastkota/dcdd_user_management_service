package services

import (
	"dcdd_user_management_service/model"

	"context"
	"io"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Services interface {
	CheckForExistingUser(field, value string) (*model.DcddUser, error)
	CreateDcddUser(signupData model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	CreateUserProfile(inputData model.UserProfileInput, tx *gorm.DB) (*model.UserProfile, error)
	UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error)
	BulkRegistration(ctx context.Context, csvData io.Reader) (error)

	FetchProfileByUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error)
}