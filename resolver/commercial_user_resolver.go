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

func (ar *UserResolver) CheckForDcddExistingUser(p graphql.ResolveParams) *model.GenericUserResponse {
	field := p.Args["field"].(string)
	value := p.Args["value"].(string)
	result, err := ar.Services.CheckForDcddExistingUser(field, value)
	if err != nil {
		return helpers.FormatError(err)
	}

	if result == nil {
        return &model.GenericUserResponse{
            Data: map[string]interface{}{
                "exist_user": false,
                "user_id":    nil,
            },
            Error: nil,
        }
    }
	
	return &model.GenericUserResponse{
		Data: map[string]interface{}{
			"exist_user": result != nil,
			"user_id":    result.ID,
		},
		Error: nil,
	}

}

func (ar *UserResolver) CreateDcddUser(p graphql.ResolveParams) *model.GenericUserResponse {

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
	return &model.GenericUserResponse{
		Data: &model.CreateUserSuccessData{
			User:    user,
			Profile: profile,
		},
		Error: nil,
	}
}

func (ur *UserResolver) CreateUserProfile(p graphql.ResolveParams) *model.GenericUserResponse {
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
    result, err := ur.Services.CreateUserProfile(userProfileInput, tx)
    if err != nil {
        tx.Rollback() // Roll back the transaction on error.
        return helpers.FormatError(err)
    }

    // Commit the transaction if everything was successful.
    if err := tx.Commit().Error; err != nil {
        return helpers.FormatError(err)
    }

    return &model.GenericUserResponse{
        Data: &model.DcddUserProfileResult{
            UserProfile: result,
        },
        Error: nil,
    }
}

func (ur *UserResolver) FetchProfileByUserId(p graphql.ResolveParams) *model.GenericUserResponse {
	userID := p.Args["user_id"].(uuid.UUID)
	result, err := ur.Services.FetchProfileByUserId(p.Context, userID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericUserResponse{
		Data: &model.DcddUserProfileResult{
			UserProfile: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) FetchAllUsers(p graphql.ResolveParams) (interface{}, error) {
	users, err := ur.Services.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserResolver) FetchAllActiveUsers(p graphql.ResolveParams)(interface{}, error) {
    users, err := ur.Services.GetAllActiveUsers()
    if err != nil {
        return nil, err
    }
    return users, nil
}
func (ur *UserResolver) FetchUsersByDateRange(p graphql.ResolveParams)(interface{}, error) {
    fromDateStr, _ := p.Args["fromDate"].(string)
    toDateStr, _ := p.Args["toDate"].(string)

    // Parse the dates (expects format YYYY-MM-DD)
    fromDate, err := time.Parse("2006-01-02", fromDateStr)
    if err != nil {
        return nil, err
    }
    toDate, err := time.Parse("2006-01-02", toDateStr)
    if err != nil {
        return nil, err
    }

    return ur.Services.FetchUsersByDateRange(fromDate, toDate)
}


func (ur *UserResolver) UpdateDcddUser(p graphql.ResolveParams) *model.GenericUserResponse {
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

    return &model.GenericUserResponse{
        Data: &model.CreateUserSuccessData{
            User:    user,
            Profile: profile,
        },
        Error: nil,
    }
}
func (ur *UserResolver) DeleteUser(p graphql.ResolveParams) *model.GenericUserResponse {
     ctx := p.Context
    userIDStr, ok := p.Args["userID"].(string)
    if !ok || userIDStr == "" {
        return helpers.FormatError(fmt.Errorf("userID argument is required"))
    }

    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        return helpers.FormatError(fmt.Errorf("invalid userID: %v", err))
    }

    updatedUser, err := ur.Services.UpdateUserStatus(ctx, userID, "Deleted")
    if err != nil {
        return helpers.FormatError(err)
    }

    return &model.GenericUserResponse{
        Data: &model.DeleteUserResult{
            User: updatedUser,
        },
        Error: nil,
    }
}


func (ur *UserResolver) UpdateUserStatus(p graphql.ResolveParams) *model.GenericUserResponse {
	userID, ok := p.Args["userID"].(uuid.UUID)
	if !ok || userID == uuid.Nil {
		return helpers.FormatError(fmt.Errorf("userID is required"))
	}

	status, ok := p.Args["status"].(string)
	if !ok {
		return helpers.FormatError(fmt.Errorf("User status is required and must be a boolean"))
	}

	result, err := ur.Services.UpdateUserStatus(p.Context, userID, status)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericUserResponse{
		Data: &model.DeleteUserResult{
			User: result,
		},
		Error: nil,
	}
}

func (ar *UserResolver) BulkRegistration(p graphql.ResolveParams) *model.GenericUserResponse {
    filePath, ok := p.Args["csv_path"].(string)
    if !ok {
        return helpers.FormatError(fmt.Errorf("csv_path argument is required"))
    }

    csvFile, err := os.Open(filePath)
    if err != nil {
        return helpers.FormatError(fmt.Errorf("failed to open file at path '%s': %w", filePath, err))
    }
    defer csvFile.Close()

    err = ar.Services.BulkRegistration(p.Context, csvFile)
    if err != nil {
        return &model.GenericUserResponse{
            Data: nil,
            Error: &model.UserError{Message: err.Error()},
        }
    }

    return &model.GenericUserResponse{
        Data: &model.BulkSuccessResult{
            Message: "Bulk registration completed successfully",
        },
        Error: nil,
    }
}


