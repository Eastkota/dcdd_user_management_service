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

func (repo *UserRepository) CheckForDcddExistingUser(field, value string) (*model.DcddUser, error) {
    var user model.DcddUser
    err := repo.DB.Where(fmt.Sprintf("%s = ? AND status != ?", field), value, "Deleted").First(&user).Error

    if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user with %v %v: %v", field, value, err)
	}
	return &user, nil
}

func (repo *UserRepository) FetchDcddUserByLoginID(field, value string) (*model.DcddUser, error) {
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
                MobileNo:       helpers.StringPtr(signupInput.MobileNo),
                Email:          helpers.StringPtr(signupInput.Email),
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
        
        userProfile, err = repo.CreateDcddUserProfile(tx, profileInput)
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

func (repo *UserRepository) CreateDcddUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error) {
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

    if inputData.GradeId != uuid.Nil && inputData.GradeId != uuid.Nil {
       userProfile.GradeId = inputData.GradeId
   }
    if inputData.SchoolId != uuid.Nil && inputData.SchoolId != uuid.Nil {
       userProfile.SchoolId = inputData.SchoolId
   }
    if inputData.EccdId != uuid.Nil && inputData.EccdId != uuid.Nil {
       userProfile.EccdId = inputData.EccdId
   }

    return &userProfile, nil
}

func (repo *UserRepository) FetchProfileByDcddUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error) {
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
func (repo *UserRepository) FetchAllDcddUsers() ([]model.DcddUserAndProfile, error) {
	var users []model.DcddUser
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
    // fmt.Println("Fetched users:", users) // <-- add this
	results := make([]model.DcddUserAndProfile, 0, len(users))

	for _, user := range users {
		var profile model.UserProfile
		err := repo.DB.Where("user_id = ?", user.ID).First(&profile).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("failed to fetch profile for user %d: %w", user.ID, err)
		}
		results = append(results, model.DcddUserAndProfile{
			User:    &user,
			UserProfile: &profile,
		})
	}

	return results, nil
}
func (repo *UserRepository) GetAllActiveDcddUsers() ([]model.DcddUserAndProfile, error){
    var users []model.DcddUser
	if err := repo.DB.Where("status = ?", "Active").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
    // fmt.Println("Fetched users:", users) // <-- add this
	results := make([]model.DcddUserAndProfile, 0, len(users))

	for _, user := range users {
		var profile model.UserProfile
		err := repo.DB.Where("user_id = ?", user.ID).First(&profile).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("failed to fetch profile for user %d: %w", user.ID, err)
		}
		results = append(results, model.DcddUserAndProfile{
			User:    &user,
			UserProfile: &profile,
		})
	}

	return results, nil
}

func (repo *UserRepository) FetchDcddUsersByDateRange(fromDate, toDate time.Time) ([]model.DcddUserAndProfile, error) {
    fromDate = fromDate.UTC()
    toDate = toDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second).UTC()

    var users []model.DcddUser
    if err := repo.DB.
        Where("created_at BETWEEN ? AND ?", fromDate, toDate).
        Find(&users).Error; err != nil {
        return nil, fmt.Errorf("failed to fetch users between %v and %v: %w", fromDate, toDate, err)
    }

    if len(users) == 0 {
        return []model.DcddUserAndProfile{}, nil
    }

    fmt.Println(len(users))

    //Collect user IDs
    userIDs := make([]uuid.UUID, len(users))
    for i, u := range users {
        userIDs[i] = u.ID
    }

    //Fetch profiles by user_id
    var profiles []model.UserProfile
    if err := repo.DB.
        Where("user_id IN ?", userIDs).
        Find(&profiles).Error; err != nil {
        return nil, fmt.Errorf("failed to fetch profiles: %w", err)
    }

    //Map profiles to their user_id
    profileMap := make(map[uuid.UUID]*model.UserProfile)
    for i, p := range profiles {
        profileMap[p.UserId] = &profiles[i]
    }

    //Combine into DcddUserAndProfile
    var results []model.DcddUserAndProfile
    for _, u := range users {
        results = append(results, model.DcddUserAndProfile{
            User:        &u,
            UserProfile: profileMap[u.ID],
        })
    }

    return results, nil
}




func (repo *UserRepository) UpdateDcddUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.DcddUser, *model.UserProfile, error) {
    var user model.DcddUser
    var userProfile model.UserProfile

    err := repo.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.First(&user, "id = ?", userID).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return fmt.Errorf("user with ID %s not found: %w", userID, err)
            }
            return fmt.Errorf("failed to fetch user for update: %w", err)
        }
        updateData := map[string]interface{}{
            "updated_at": time.Now(),
        }

        if signupInput.Password != "" {
            hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
            if err != nil {
                return fmt.Errorf("failed to hash password: %w", err)
            }
            updateData["password"] = hashedPassword
        }
        if signupInput.MobileNo != "" {
            updateData["mobile_no"] = signupInput.MobileNo
        }
        if signupInput.Email != "" {
            updateData["email"] = signupInput.Email
        }

        if err := tx.Model(&user).Updates(updateData).Error; err != nil {
            return fmt.Errorf("failed to update user data: %w", err)
        }

        if err := tx.Where("user_id = ?", user.ID).First(&userProfile).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return fmt.Errorf("user profile not found for user ID %s: %w", user.ID, err)
            }
            return fmt.Errorf("failed to fetch user profile: %w", err)
        }

        profileUpdate := map[string]interface{}{}
        if signupInput.Name != "" {
            profileUpdate["name"] = signupInput.Name
        }
        if signupInput.Gender != "" {
            profileUpdate["gender"] = signupInput.Gender
        }
        if signupInput.Cid != "" {
            profileUpdate["cid"] = signupInput.Cid
        }
        if len(profileUpdate) > 0 {
            if err := tx.Model(&userProfile).Updates(profileUpdate).Error; err != nil {
                return fmt.Errorf("failed to update user profile: %w", err)
            }
        }

        return nil
    })

    if err != nil {
        return nil, nil, err
    }

    return &user, &userProfile, nil
}

func (repo *UserRepository) UpdateDcddUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.DcddUser, error) {
	// Find the existing question by its ID
	var user model.DcddUser
	if err := repo.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	user.Status = status
	if err := repo.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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
                MobileNo:       helpers.StringPtr(signupInput.MobileNo),
                Email:          helpers.StringPtr(signupInput.Email),
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
            userProfile, err := repo.CreateDcddUserProfile(tx, profileInput)
            if err != nil {
                return fmt.Errorf("failed to create user profile for user '%s': %v", err)
            }
            userProfiles = append(userProfiles, userProfile)
        }
        return nil
    })
    return err
}

func (repo *UserRepository) GetSchool(ctx context.Context, id uuid.UUID) (*model.School, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var school model.School
	if err := db.WithContext(ctx).Where("id = ?", id).First(&school).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch school: %v", err)
	}
	return &school, nil
}


func (repo *UserRepository) GetEccd(ctx context.Context, id uuid.UUID) (*model.Eccd, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var eccd model.Eccd
	if err := db.WithContext(ctx).Where("id = ?", id).First(&eccd).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch eccd: %v", err)
	}
	return &eccd, nil
}

func (repo *UserRepository) GetGrade(ctx context.Context, id uuid.UUID) (*model.Grade, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var grade model.Grade
	if err := db.WithContext(ctx).Where("id = ?", id).First(&grade).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch grade: %v", err)
	}
	return &grade, nil
}

func (repo *UserRepository) GetDzongkhag(ctx context.Context, id uuid.UUID) (*model.Dzongkhag, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var dzongkhag model.Dzongkhag
	if err := db.WithContext(ctx).Where("id = ?", id).First(&dzongkhag).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch dzongkhag: %v", err)
	}
	return &dzongkhag, nil
}

func (repo *UserRepository) FetchDzongkhag(ctx context.Context) ([]model.Dzongkhag, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var dzongkhags []model.Dzongkhag
	if err := db.WithContext(ctx).Find(&dzongkhags).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch dzongkhags: %v", err)
	}
	return dzongkhags, nil
}

func (repo *UserRepository) FetchGrade(ctx context.Context) ([]model.Grade, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var grades []model.Grade
	if err := db.WithContext(ctx).Find(&grades).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch grades: %v", err)
	}
	return grades, nil
}

func (repo *UserRepository) FetchSchool(ctx context.Context, dzongkhagId uuid.UUID) ([]model.School, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }
	var schools []model.School
	if err := db.WithContext(ctx).Where("dzongkhag_id = ? ", dzongkhagId).Find(&schools).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch schools: %v", err)
	}
	return schools, nil
}


func (repo *UserRepository) FetchEccd(ctx context.Context, DzonghkhagId uuid.UUID) ([]model.Eccd, error) {
	db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

	var eccds []model.Eccd
	if err := db.WithContext(ctx).Where("dzongkhag_id = ? ", DzonghkhagId).Find(&eccds).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch eccds: %v", err)
	}
	return eccds, nil
}

