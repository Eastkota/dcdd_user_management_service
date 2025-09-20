package services

import (
	"dcdd_user_management_service/repositories"
	"dcdd_user_management_service/model"
	"dcdd_user_management_service/helpers"

    "time"
	"fmt"
	"context"
	"io"
	"encoding/csv"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UserService struct {
	Repository repositories.Repository // Inject Repository
}

func NewUserService(repository repositories.Repository) *UserService {
	return &UserService{Repository: repository}
}

func (as *UserService) CheckForDcddExistingUser(field, value string) (*model.DcddUser, error) {
	return as.Repository.CheckForDcddExistingUser(field, value)
}

func (as *UserService) CreateDcddUser(signupData model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {

	if signupData.MobileNo == "" && signupData.Email == "" {
		return nil, nil, fmt.Errorf("either email or mobile number is required")
	}
	if signupData.MobileNo != "" {
		mobileResult, err := as.CheckForDcddExistingUser("mobile_no", signupData.MobileNo)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to register: %v", err)
		}

		if mobileResult != nil {
			return nil, nil, fmt.Errorf("user with given mobile number %v already exist", signupData.MobileNo)
		}
	}

	if signupData.Email != "" {
		emailResult, err := as.CheckForDcddExistingUser("email", signupData.Email)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to register %v", err)
		}
		if emailResult != nil {
			return nil, nil, fmt.Errorf("user with given email address %v already exist", signupData.Email)
		}
	}

	return as.Repository.CreateDcddUser(&signupData)
}

func (as *UserService) CreateDcddUserProfile(inputData model.UserProfileInput, tx *gorm.DB) (*model.UserProfile, error) {
    // Pass the transaction 'tx' to the repository method.
    return as.Repository.CreateDcddUserProfile(tx, inputData)
}

func (as *UserService) UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {
	return as.Repository.UpdateDcddUser(userID, signupInput)
}

func (as *UserService) FetchProfileByDcddUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error) {
    return as.Repository.FetchProfileByDcddUserId(ctx, userID)
}

func (as *UserService) UpdateDcddUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error) {
	return as.Repository.UpdateDcddUserStatus(ctx, userID, status)
}

func (as *UserService) GetAllDcddUsers() ([]model.DcddUser, error) {
	return as.Repository.GetAllDcddUsers()
}

func (as *UserService) GetAllActiveDcddUsers() ([]model.DcddUser, error) {
    return as.Repository.GetAllActiveDcddUsers()
}
func (as *UserService) FetchDcddUsersByDateRange(fromDate, toDate time.Time) ([]model.DcddUser, error) {
    return as.Repository.FetchDcddUsersByDateRange(fromDate, toDate)
}

func (as *UserService) BulkRegistration(ctx context.Context, csvData io.Reader) error {
    reader := csv.NewReader(csvData)

    records, err := reader.ReadAll()
    if err != nil {
        return fmt.Errorf("failed to read CSV file: %w", err)
    }

    if len(records) == 0 || !helpers.ValidateCSVHeader(records[0]) {
        return fmt.Errorf("invalid CSV header. Expected 'Name,MobileNo,Email,Gender,Password,StudentId,Category,SchoolId,GradeId,EccdId,Dob,DzongkhagId,Cid'")
    }
    records = records[1:]

    var signupInputs []model.SignupInput
    for i, record := range records {
        if len(record) < 13 {
            return fmt.Errorf("invalid record format at row %d: row has too few columns (expected at least 13, got %d)", i+2, len(record))
        }

        // --- Perform Type Conversions Here ---
        
        // Convert SchoolId, GradeId, EccdId, and DzongkhagId from string to uuid.NullUUID
        schoolID, err := helpers.ParseUUIDStringToNullUUID(record[7])
        if err != nil {
            return fmt.Errorf("failed to parse SchoolId at row %d: %w", i+2, err)
        }

        gradeID, err := helpers.ParseUUIDStringToNullUUID(record[8])
        if err != nil {
            return fmt.Errorf("failed to parse GradeId at row %d: %w", i+2, err)
        }
        
        eccdID, err := helpers.ParseUUIDStringToNullUUID(record[9])
        if err != nil {
            return fmt.Errorf("failed to parse EccdId at row %d: %w", i+2, err)
        }

        dzongkhagID, err := helpers.ParseUUIDStringToNullUUID(record[11])
        if err != nil {
            return fmt.Errorf("failed to parse DzongkhagId at row %d: %w", i+2, err)
        }

        // Convert Dob (Date of Birth) from string to *time.Time
        dob, err := helpers.ParseDateString(record[10])
        if err != nil {
            return fmt.Errorf("failed to parse Dob at row %d: %w", i+2, err)
        }

        signupInput := model.SignupInput{
            Name:        record[0],
            MobileNo:    record[1],
            Email:       record[2],
            Gender:      record[3],
            Password:    record[4],
            StudentId:   record[5],
            Category:    record[6],
            SchoolId:    schoolID.UUID,      
            GradeId:     gradeID.UUID,       
            EccdId:      eccdID.UUID,        
            Dob:         dob,           
            DzongkhagId: dzongkhagID.UUID,   
            Cid:         record[12],
        }

        if signupInput.MobileNo == "" && signupInput.Email == "" {
            return fmt.Errorf("either email or mobile number is required for user '%s' at row %d", signupInput.Name, i+2)
        }

        if signupInput.MobileNo != "" {
            if exists, err := as.CheckForDcddExistingUser("mobile_no", signupInput.MobileNo); err != nil || exists != nil {
                return fmt.Errorf("user with mobile number '%s' at row %d already exists", signupInput.MobileNo, i+2)
            }
        }
        if signupInput.Email != "" {
            if exists, err := as.CheckForDcddExistingUser("email", signupInput.Email); err != nil || exists != nil {
                return fmt.Errorf("user with email '%s' at row %d already exists", signupInput.Email, i+2)
            }
        }

        signupInputs = append(signupInputs, signupInput)
    }

    return as.Repository.BulkRegistration(signupInputs)
}