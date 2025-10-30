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
	UpdateDcddUserPassword(userID uuid.UUID, DcddResetPasswordInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error)
	GetAllDcddUsers(limit, offset int) ([]model.DcddUser, int, error)
	GetAllActiveDcddUsers(limit, offset int) ([]model.DcddUser, int, error)
	FetchDcddUsersByDateRange(fromDate, toDate time.Time) ([]model.DcddUserAndProfile, error)
	UpdateDcddUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error)
	BulkRegistration(ctx context.Context, csvData io.Reader) (error)

	FetchProfileByDcddUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error)
	FetchDzongkhag(ctx context.Context) ([]model.Dzongkhag, error)
	FetchGrade(ctx context.Context) ([]model.Grade, error)
	FetchSchool(ctx context.Context, schoolId uuid.UUID) ([]model.School, error)
	FetchEccd(ctx context.Context, schoolId  uuid.UUID) ([]model.Eccd, error)
	GetDcddUserTotals(fromDate, toDate *time.Time) (totalAll int, totalActive int, totalNew int, err error)

	GetUserActivity(offset, limit int) ([]model.UserActivity, error)
}