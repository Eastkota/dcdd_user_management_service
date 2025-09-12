package repositories

import (
    "dcdd_user_management_service/model"
    "dcdd_user_management_service/helpers"

    "fmt"
    "time"
    "errors"
    "context"

    "gorm.io/gorm"
	"github.com/google/uuid"
)

type UserRepository struct{
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (repo *UserRepository) CheckForExistingUser(field, value string) (*model.DcddUser, error) {
    var user model.DcddUser
    err := repo.DB.Where(fmt.Sprintf("%s = ? AND status != ?", field), value, "Deleted").First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("failed to find user with %v %v", field, value)
    }
    return &user, nil
}

func (repo *UserRepository) FetchUserByLoginID(field, value string) (*model.DcddUser, error) {
    var user model.DcddUser
    err := repo.DB.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error
    if err != nil {
        return nil, fmt.Errorf("failed to find user with %v %v", field, value)
    }
    return &user, nil
}

func (repo *UserRepository) CreateDcddUser(signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {
    var user *model.DcddUser
    var createdUser *model.DcddUser
    var userProfile *model.UserProfile
    var err error

    err = repo.DB.Transaction(func(tx *gorm.DB) error {
        identifier, err := helpers.GenerateRandomTokenString(6)
        if err != nil {
            return fmt.Errorf("failed to generate identifier: %v", err)
        }

        hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
        if err != nil {
            return fmt.Errorf("failed to hash password: %v", err)
        }

        // Check for existing user by mobile and email within the transaction.
        var existingUser model.DcddUser
        if signupInput.MobileNo != "" {
            tx.Where("mobile_no = ?", signupInput.MobileNo).First(&existingUser)
        }
        if existingUser.ID == uuid.Nil && signupInput.Email != "" {
            tx.Where("email = ?", signupInput.Email).First(&existingUser)
        }
        user = &existingUser

        loginid := helpers.GenerateLoginId(6)

        // If a user is found, update their record.
        if user.ID != uuid.Nil {
            updateData := map[string]interface{}{
                "mobile_no":       signupInput.MobileNo,
                "email":           signupInput.Email,
                "user_identifier": identifier,
                "password":        hashedPassword,
                "status":          "Active",
                "updated_at":      time.Now(),
                "category":        signupInput.Category,
                "login_id":        loginid,
            }
            if err := tx.Model(user).Updates(updateData).Error; err != nil {
                return fmt.Errorf("failed to update user data: %v", err)
            }
            createdUser = user
            
        } else {
            // If no user is found, create a new one.
            newUser := model.DcddUser{
                ID:             uuid.New(),
                MobileNo:       signupInput.MobileNo,
                Email:          signupInput.Email,
                UserIdentifier: identifier,
                Password:       hashedPassword,
                Status:         "Active",
                CreatedAt:      time.Now(),
                UpdatedAt:      time.Now(),
                StudentId :     signupInput.StudentId,
                Category :      signupInput.Category,
                LoginId :       loginid,    
            }
            if err := tx.Create(&newUser).Error; err != nil {
                return fmt.Errorf("failed to insert user data: %v", err)
            }
            createdUser = &newUser
        }

        profileInput := model.UserProfileInput{
            UserId:     createdUser.ID,
            Gender:     signupInput.Gender,
            Name:       signupInput.Name,
            Cid:        signupInput.Cid,
            Dob:        signupInput.Dob,
            DzongkhagId: signupInput.DzongkhagId,
            EccdId:     signupInput.EccdId,
            SchoolId:   signupInput.SchoolId,
            GradeId:    signupInput.GradeId,
        }
        
        userProfile, err = repo.CreateUserProfile(tx, profileInput)
        if err != nil {
            return fmt.Errorf("failed to create user profile: %v", err)
        }
        
        return nil
    })

    if err != nil {
        return nil, nil, err
    }

    return createdUser, userProfile, nil
}

func (repo *UserRepository) CreateUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error) {
    userProfile := model.UserProfile{
        ID:             uuid.New(),
        Name:           inputData.Name,
        Gender:         inputData.Gender,
        ProfilePicture: inputData.ProfilePicture,
        UserId:         inputData.UserId,
        GradeId:        inputData.GradeId,
        SchoolId:       inputData.SchoolId,
        EccdId:         inputData.EccdId,
        DzongkhagId:    inputData.DzongkhagId,
        Dob:            inputData.Dob,
        Cid:            inputData.Cid,
        CreatedAt:      time.Now(),
        UpdatedAt:      time.Now(),
    }

    if err := tx.Create(&userProfile).Error; err != nil {
        return nil, fmt.Errorf("failed to insert user profile: %v", err)
    }


    // profile, err := repo.FetchProfileByUserId(context.Background(), inputData.UserId)
    // if err != nil {
    //     return nil, fmt.Errorf("failed to fetch created profile: %v", err)
    // }
    return &userProfile, nil
}

func (repo *UserRepository) FetchProfileByUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error) {
    var profile model.UserProfile

    db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

    if err := db.WithContext(ctx).Where("user_id = ?", userId).First(&profile).Error; err != nil {
        return nil, fmt.Errorf("failed to fetch profile for user id '%s': %w", userId, err)
    }
    return &profile, nil
}

func (repo *UserRepository) UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {
    var user model.DcddUser
    var userProfile *model.UserProfile
    var err error

    err = repo.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.First(&user, "id = ?", userID).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return fmt.Errorf("user with ID %s not found: %w", userID, err)
            }
            return fmt.Errorf("failed to fetch user for update: %w", err)
        }

        // Step 2: Prepare the update data.
        hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
        if err != nil {
            return fmt.Errorf("failed to hash password: %w", err)
        }

        updateData := map[string]interface{}{
            "name":       signupInput.Name,
            "password":   hashedPassword,
            "updated_at": time.Now(),
        }
        
        if signupInput.MobileNo != "" {
            updateData["mobile_no"] = signupInput.MobileNo
        }
        if signupInput.Email != "" {
            updateData["email"] = signupInput.Email
        }

        if err := tx.Model(&model.DcddUser{}).Where("id = ?", userID).Updates(updateData).Error; err != nil {
            return fmt.Errorf("failed to update user data: %w", err)
        }

        // The 'user' variable is from the tx.First call, so we can use its ID.
        userProfile, err = repo.FetchProfileByUserId(context.Background(), user.ID)
        if err != nil {
            return fmt.Errorf("failed to fetch user profile after update: %w", err)
        }

        return nil
    })

    if err != nil {
        return nil, nil, err
    }


    var updatedUser model.DcddUser
    if err := repo.DB.First(&updatedUser, "id = ?", userID).Error; err != nil {
        return nil, nil, fmt.Errorf("failed to re-fetch updated user: %w", err)
    }

    return &updatedUser, userProfile, nil
}

func (repo *UserRepository) UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error) {
	// Find the existing question by its ID
	var user model.DcddUser
	if err := repo.DB.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("user not found with ID: %s", userID)
	}

	updates := map[string]interface{}{
		"status": status,
		"updated_at":   time.Now(),
	}

	if err := repo.DB.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update user status: %w", err)
	}

	// Fetch the updated question to return the new state.
	var updatedUser model.DcddUser
	if err := repo.DB.WithContext(ctx).Where("id = ?", userID).First(&updatedUser).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch the updated USER: %w", err)
	}

	return &updatedUser, nil
}

func (repo *UserRepository) BulkRegistration(signupInputs []model.SignupInput) (error) {
    var users []*model.DcddUser
    var userProfiles []*model.UserProfile

    // Start a single transaction for the entire batch.
    err := repo.DB.Transaction(func(tx *gorm.DB) error {
        for _, signupInput := range signupInputs {
            identifier, err := helpers.GenerateRandomTokenString(6)
            if err != nil {
                return fmt.Errorf("failed to generate identifier: %v", err)
            }

            hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
            if err != nil {
                return fmt.Errorf("failed to hash password: %v", err)
            }

            loginid := helpers.GenerateLoginId(6)

            newUser := model.DcddUser{
                ID:             uuid.New(),
                MobileNo:       signupInput.MobileNo,
                Email:          signupInput.Email,
                UserIdentifier: identifier,
                Password:       hashedPassword,
                Status:         "Active",
                CreatedAt:      time.Now(),
                UpdatedAt:      time.Now(),
                StudentId:     signupInput.StudentId,
                Category:      signupInput.Category,
                LoginId:       loginid,
            }

            if err := tx.Create(&newUser).Error; err != nil {
                return fmt.Errorf("failed to insert user data for '%s': %v", signupInput.Name, err)
            }
            users = append(users, &newUser)

            profileInput := model.UserProfileInput{
                UserId: newUser.ID,
                Gender: signupInput.Gender,
                Name:   signupInput.Name,
                MobileNo: signupInput.MobileNo,
                Email:    signupInput.Email,
                GradeId:  signupInput.GradeId,
                StudentId: signupInput.StudentId,
                Password:  signupInput.Password,
                SchoolId:  signupInput.SchoolId,
                EccdId:    signupInput.EccdId,
                Dob:       signupInput.Dob,
                Cid:       signupInput.Cid,
                DzongkhagId: signupInput.DzongkhagId,
                Category:signupInput.Category,
            }
            userProfile, err := repo.CreateUserProfile(tx, profileInput)
            if err != nil {
                return fmt.Errorf("failed to create user profile for user '%s': %v", err)
            }
            userProfiles = append(userProfiles, userProfile)
        }
        return nil
    })
    return err
}

