package resolvers

import (
	"dcdd_user_management_service/helpers"
	"dcdd_user_management_service/model"
	"dcdd_user_management_service/services"

	"os"
	"fmt"
    "time"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/google/uuid"
)

type UserResolver struct {
	Services services.Services // Inject Services
}

func NewUserResolver(service services.Services) *UserResolver {
	return &UserResolver{Services: service}
}

func (ar *UserResolver) CheckForDcddExistingUser(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	field := p.Args["field"].(string)
	value := p.Args["value"].(string)
	result, err := ar.Services.CheckForDcddExistingUser(field, value)
	if err != nil {
		return helpers.FormatError(err)
	}

	if result == nil {
        return &model.DcddGenericUserResponse{
            Data: map[string]interface{}{
                "exist_user": false,
                "user_id":    nil,
            },
            Error: nil,
        }
    }
	return &model.DcddGenericUserResponse{
		Data: map[string]interface{}{
			"exist_user": result != nil,
			"user_id":    result.ID,
		},
		Error: nil,
	}

}

func (ar *UserResolver) CreateDcddUser(p graphql.ResolveParams) *model.DcddGenericUserResponse {

	var signupInput model.SignupInput
	inputData := p.Args["signup_input"].(map[string]interface{})

	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = json.Unmarshal(jsonData, &signupInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	user, profile, err := ar.Services.CreateDcddUser(signupInput)
	if err != nil {
		return helpers.FormatError(err)
	}
    fmt.Println("Returning user profile:", user, profile)

	return &model.DcddGenericUserResponse{
		Data: &model.CreateUserSuccessData{
			User:    user,
			Profile: profile,
		},
		Error: nil,
	}
}

func (ur *UserResolver) CreateDcddUserProfile(p graphql.ResolveParams) *model.DcddGenericUserResponse {
    var userProfileInput model.UserProfileInput
    inputData := p.Args["input"].(map[string]interface{})
    jsonData, err := json.Marshal(inputData)
    if err != nil {
        return helpers.FormatError(err)
    }
    err = json.Unmarshal(jsonData, &userProfileInput)
    if err != nil {
        return helpers.FormatError(err)
    }

    // Get the database connection from the context or a global variable.
    // Assuming 'ur.DB' holds the GORM database instance.
    db, err := helpers.GetGormDB()

    // Start a new transaction.
    tx := db.Begin()
    if tx.Error != nil {
        return helpers.FormatError(tx.Error)
    }

    // Defer a rollback in case of an error.
    // If the function returns successfully, the commit will prevent this.
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Pass the transaction 'tx' to the service layer.
    result, err := ur.Services.CreateDcddUserProfile(userProfileInput, tx)
    if err != nil {
        tx.Rollback() // Roll back the transaction on error.
        return helpers.FormatError(err)
    }

    // Commit the transaction if everything was successful.
    if err := tx.Commit().Error; err != nil {
        return helpers.FormatError(err)
    }

    return &model.DcddGenericUserResponse{
        Data: &model.DcddUserProfileResult{
            UserProfile: result,
        },
        Error: nil,
    }
}

func (ur *UserResolver) FetchProfileByDcddUserId(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	userID := p.Args["user_id"].(uuid.UUID)
	result, err := ur.Services.FetchProfileByDcddUserId(p.Context, userID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.DcddGenericUserResponse{
		Data: &model.DcddUserProfileResult{
			UserProfile: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) FetchAllUsers(p graphql.ResolveParams) (interface{}, error) {
	users, err := ur.Services.GetAllDcddUsers()
    if err != nil {
        return map[string]interface{}{
            "data":  nil,
            "error": err.Error(),
        }, nil
    }

    return map[string]interface{}{
        "data":  users,
        "error": nil,
    }, nil
}

func (ur *UserResolver) FetchAllActiveUsers(p graphql.ResolveParams)(interface{}, error) {
    users, err := ur.Services.GetAllActiveDcddUsers()
    if err != nil {
        return map[string]interface{}{
            "data":  nil,
            "error": err.Error(),
        }, nil
    }

    return map[string]interface{}{
        "data":  users,
        "error": nil,
    }, nil
}

func (ur *UserResolver) FetchDcddUsersByDateRange(p graphql.ResolveParams) (interface{}, error) {
    fromDateStr, _ := p.Args["fromDate"].(string)
    toDateStr, _ := p.Args["toDate"].(string)

    fromDate, err := time.Parse("2006-01-02", fromDateStr)
    if err != nil {
        return map[string]interface{}{
            "data":  nil,
            "error": err.Error(),
        }, nil
    }
    toDate, err := time.Parse("2006-01-02", toDateStr)
    if err != nil {
        return map[string]interface{}{
            "data":  nil,
            "error": err.Error(),
        }, nil
    }

    users, err := ur.Services.FetchDcddUsersByDateRange(fromDate, toDate)
    if err != nil {
        return map[string]interface{}{
            "data":  nil,
            "error": err.Error(),
        }, nil
    }

    return map[string]interface{}{
        "data":  users,
        "error": nil,
    }, nil
}


func (ur *UserResolver) UpdateDcddUser(p graphql.ResolveParams) *model.DcddGenericUserResponse {
    var signupInput model.SignupInput
    userID, ok := p.Args["user_id"].(uuid.UUID)
    if !ok {
        return helpers.FormatError(fmt.Errorf("user_id argument is not a valid UUID type"))
    }

    inputData, ok := p.Args["signup_input"].(map[string]interface{})
    if !ok {
        return helpers.FormatError(fmt.Errorf("signup_input argument is not a valid map"))
    }

    jsonData, err := json.Marshal(inputData)
    if err != nil {
        return helpers.FormatError(err)
    }

    err = json.Unmarshal(jsonData, &signupInput)
    if err != nil {
        return helpers.FormatError(err)
    }

    user, profile, err := ur.Services.UpdateDcddUser(userID, &signupInput)
    if err != nil {
        return helpers.FormatError(err)
    }

    return &model.DcddGenericUserResponse{
        Data: &model.CreateUserSuccessData{
            User:    user,
            Profile: profile,
        },
        Error: nil,
    }
}
func (ur *UserResolver) DeleteUser(p graphql.ResolveParams) *model.DcddGenericUserResponse {
     ctx := p.Context
    userIDStr, ok := p.Args["userID"].(string)
    if !ok || userIDStr == "" {
        return helpers.FormatError(fmt.Errorf("userID argument is required"))
    }

    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        return helpers.FormatError(fmt.Errorf("invalid userID: %v", err))
    }

    updatedUser, err := ur.Services.UpdateDcddUserStatus(ctx, userID, "Deleted")
    if err != nil {
        return helpers.FormatError(err)
    }

    return &model.DcddGenericUserResponse{
        Data: &model.DeleteUserResult{
            User: updatedUser,
        },
        Error: nil,
    }
}


func (ur *UserResolver) UpdateDcddUserStatus(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	userID, ok := p.Args["userID"].(uuid.UUID)
	if !ok || userID == uuid.Nil {
		return helpers.FormatError(fmt.Errorf("userID is required"))
	}

	status, ok := p.Args["status"].(string)
	if !ok {
		return helpers.FormatError(fmt.Errorf("User status is required and must be a boolean"))
	}

	result, err := ur.Services.UpdateDcddUserStatus(p.Context, userID, status)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.DcddGenericUserResponse{
		Data: &model.DeleteUserResult{
			User: result,
		},
		Error: nil,
	}
}

func (ar *UserResolver) BulkRegistration(p graphql.ResolveParams) *model.DcddGenericUserResponse {
    filePath, ok := p.Args["csv_path"].(string)
    if !ok {
        return helpers.FormatError(fmt.Errorf("csv_path argument is required"))
    }
    
    // Adjust path for file system access
    fileSystemPath := filePath
    if len(filePath) > 0 && filePath[0] == '/' {
        fileSystemPath = filePath[1:] 
    }

    csvFile, err := os.Open(fileSystemPath)
    if err != nil {
        return helpers.FormatError(fmt.Errorf("failed to open file at path '%s': %w", fileSystemPath, err))
    }
    
    // Defer closing the file
    defer csvFile.Close() 
    defer func() {
        if r := os.Remove(fileSystemPath); r != nil {
            fmt.Printf("Warning: Failed to delete CSV file '%s': %v\n", fileSystemPath, r)
        } else {
            fmt.Printf("Successfully deleted CSV file: %s\n", fileSystemPath)
        }
    }()
    
    // Call the service function to perform the registration
    err = ar.Services.BulkRegistration(p.Context, csvFile)
    
    if err != nil {
        return &model.DcddGenericUserResponse{
            Data: nil,
            Error: &model.UserError{Message: err.Error()},
        }
    }

    return &model.DcddGenericUserResponse{
        Data: &model.BulkSuccessResult{
            Message: "Bulk registration completed successfully and file deleted.", // Updated message
        },
        Error: nil,
    }
}

func (ur *UserResolver) FetchDzongkhags(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	result, err := ur.Services.FetchDzongkhag(p.Context)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.DcddGenericUserResponse{
		Data: &model.DzongkhagResult{
			Dzongkhags: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) FetchSchool(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	dzongkhagId := p.Args["dzongkhag_id"].(uuid.UUID)
	result, err := ur.Services.FetchSchool(p.Context, dzongkhagId)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.DcddGenericUserResponse{
		Data: &model.SchoolResult{
			School: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) FetchEccd(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	dzongkhagId := p.Args["dzongkhag_id"].(uuid.UUID)
	result, err := ur.Services.FetchEccd(p.Context, dzongkhagId)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.DcddGenericUserResponse{
		Data: &model.EccdResult{
			Eccd: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) FetchGrade(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	result, err := ur.Services.FetchGrade(p.Context)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.DcddGenericUserResponse{
		Data: &model.GradeResult{
			Grades: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) FetchDzongkhag(p graphql.ResolveParams) *model.DcddGenericUserResponse {
	result, err := ur.Services.FetchDzongkhag(p.Context)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.DcddGenericUserResponse{
		Data: &model.DzongkhagResult{
			Dzongkhags: result,
		},
		Error: nil,
	}
}

