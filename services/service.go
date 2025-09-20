package services

import (
	"dcdd_user_management_service/model"
	
	"io"
	"time"
	"context"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Services interface {
	CheckForDcddExistingUser(field, value string) (*model.DcddUser, error)
	CreateDcddUser(signupData model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	CreateDcddUserProfile(inputData model.UserProfileInput, tx *gorm.DB) (*model.UserProfile, error)
	UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	GetAllDcddUsers() ([]model.DcddUser, error)
	GetAllActiveDcddUsers() ([]model.DcddUser, error)
	FetchDcddUsersByDateRange(fromDate, toDate time.Time) ([]model.DcddUser, error)
	UpdateDcddUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error)
	BulkRegistration(ctx context.Context, csvData io.Reader) (error)

	FetchProfileByDcddUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error)
}