package services

import (
	"user_management_service/repositories"
	"user_management_service/model"

	"fmt"
	"context"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UserService struct {
	Repository repositories.Repository // Inject Repository
}

func NewUserService(repository repositories.Repository) *UserService {
	return &UserService{Repository: repository}
}

func (as *UserService) CheckForExistingUser(field, value string) (*model.DcddUser, error) {
	return as.Repository.CheckForExistingUser(field, value)
}

func (as *UserService) CreateDcddUser(signupData model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {

	if signupData.MobileNo == "" && signupData.Email == "" {
		return nil, nil, fmt.Errorf("either email or mobile number is required")
	}
	if signupData.MobileNo != "" {
		mobileResult, err := as.CheckForExistingUser("mobile_no", signupData.MobileNo)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to register: %v", err)
		}

		if mobileResult != nil {
			return nil, nil, fmt.Errorf("user with given mobile number %v already exist", signupData.MobileNo)
		}
	}

	if signupData.Email != "" {
		emailResult, err := as.CheckForExistingUser("email", signupData.Email)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to register %v", err)
		}
		if emailResult != nil {
			return nil, nil, fmt.Errorf("user with given email address %v already exist", signupData.Email)
		}
	}

	return as.Repository.CreateDcddUser(&signupData)
}

func (as *UserService) CreateUserProfile(inputData model.UserProfileInput, tx *gorm.DB) (*model.UserProfile, error) {
    // Pass the transaction 'tx' to the repository method.
    return as.Repository.CreateUserProfile(tx, inputData)
}

func (as *UserService) UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {
	return as.Repository.UpdateDcddUser(userID, signupInput)
}

func (as *UserService) FetchProfileByUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error) {
    return as.Repository.FetchProfileByUserId(ctx, userID)
}

func (vs *UserService) UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error) {
	return vs.Repository.UpdateUserStatus(ctx, userID, status)
}