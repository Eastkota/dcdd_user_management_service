package repositories

import (
	"user_management_service/model"

	"context"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	CheckForExistingUser(field, value string) (*model.DcddUser, error)
	CreateDcddUser(signUpInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	FetchUserByLoginID(field, value string) (*model.DcddUser, error)
	CreateUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error)
	UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	FetchProfileByUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error)
	UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error)

	FetchDzongkhag(ctx context.Context) ([]model.Dzongkhag, error)
	FetchGrade(ctx context.Context) ([]model.Grade, error)
	FetchSchool(ctx context.Context, dzongkhagId uuid.UUID) ([]model.School, error)
	FetchEccd(ctx context.Context, dzongkhagId uuid.UUID) ([]model.Eccd, error)
}
