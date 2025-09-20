package repositories

import (
	"dcdd_user_management_service/model"
	"time"
	"context"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	CheckForDcddExistingUser(field, value string) (*model.DcddUser, error)
	CreateDcddUser(signUpInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	CreateDcddUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error)
	UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	GetAllDcddUsers() ([]model.DcddUser, error)
	GetAllActiveDcddUsers() ([]model.DcddUser, error)
	FetchDcddUsersByDateRange(fromDate, toDate time.Time) ([]model.DcddUser, error)
	UpdateDcddUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error)
	BulkRegistration(signupInputs []model.SignupInput) (error) 

	FetchProfileByDcddUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error)
	FetchDcddUserByLoginID(field, value string) (*model.DcddUser, error)
}
