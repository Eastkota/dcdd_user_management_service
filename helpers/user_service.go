package helpers

import (
	"dcdd_user_management_service/config"
	"dcdd_user_management_service/model"

    "context"
    "fmt"
    "time"
	"github.com/machinebox/graphql"
    "github.com/google/uuid"
)

func CreateDcddUserProfile(profileInputData map[string]interface{}) (*model.UserProfile, error) {
    cleanData := CleanInterfaceData(profileInputData)
	userServiceClient := graphql.NewClient(config.UserServiceApiUrl())

	req := graphql.NewRequest(`
		mutation DcddCreateUserProfile($input: DcddUserProfileInput) {
            dcddCreateUserProfile(input: $input) {
                data {
                    user_profile {
                        cid
                        created_at
                        dob
                        dzongkhag_id
                        eccd_id
                        gender
                        grade_id
                        id
                        name
                        profile_picture
                        school_id
                        updated_at
                        user_id
                        dzongkhag {
                            id
                            name
                        }
                        eccd {
                            id
                            name
                            sort
                        }
                        grade {
                            id
                            name
                        }
                        school {
                            id
                            pvt_public
                            name
                        }
                    }
                }
                error {
                    code
                    field
                    message
                }
            }
        }`)

	req.Var("input", cleanData)
	req.Header.Set("Cache-Control", "no-cache")
	var tempResponse struct {
        DcddCreateUserProfile struct {
            Data struct {
                UserProfile struct {
                    DzongkhagId uuid.UUID   `json:"dzongkhag_id"` 
                    ID         uuid.UUID `json:"id"`
                    Gender    string   `json:"gender"`
                    EccdId    uuid.UUID   `json:"eccd_id"`
                    Cid       string   `json:"cid"`
                    Name      string   `json:"name"`
                    UpdatedAt time.Time   `json:"updated_at"`
                    UserId    uuid.UUID   `json:"user_id"`
                    GradeId   uuid.UUID   `json:"grade_id"`
                    SchoolId  uuid.UUID   `json:"school_id"`
                    ProfilePicture string   `json:"profile_picture"`
                    CreatedAt time.Time   `json:"created_at"`
                    Dob     *time.Time   `json:"dob"`
                    Dzongkhag struct {
                        ID   uuid.UUID `json:"id"`
                        Name string `json:"name"`
                    } `json:"dzongkhag"`
                    Eccd struct {
                        ID        uuid.UUID `json:"id"`
                        Name      string `json:"name"`
                        Sort      string `json:"sort"`
                    } `json:"eccd"`
                    Grade struct {
                        ID   uuid.UUID `json:"id"`
                        Name string `json:"name"`
                    } `json:"grade"`
                    School struct {
                        ID        uuid.UUID `json:"id"`
                        PvtPublic string `json:"pvt_public"`
                        Name      string `json:"name"`  
                    } `json:"school"`           
                } `json:"user_profile"`
            } `json:"data"`
            Error struct {
                Code    string `json:"code"`
                Field   string `json:"field"`
                Message string `json:"message"`
            } `json:"error"`
        } `json:"dcddCreateUserProfile"`
    }

	// Execute the request
	err := userServiceClient.Run(context.Background(), req, &tempResponse)
    fmt.Println("Response from user service:", err, tempResponse)
	if err != nil {
		return nil, err
	}
	// Check for errors in response
	if tempResponse.DcddCreateUserProfile.Error.Message != "" {
		return nil, fmt.Errorf("error from video service: %s", tempResponse.DcddCreateUserProfile.Error.Message)
	}


	formattedProfile := &model.UserProfile{
        DzongkhagId:    tempResponse.DcddCreateUserProfile.Data.UserProfile.DzongkhagId,
        ID:             tempResponse.DcddCreateUserProfile.Data.UserProfile.ID,
        Gender:         tempResponse.DcddCreateUserProfile.Data.UserProfile.Gender,
        EccdId:         tempResponse.DcddCreateUserProfile.Data.UserProfile.EccdId,
        Cid:            tempResponse.DcddCreateUserProfile.Data.UserProfile.Cid,
        Name:           tempResponse.DcddCreateUserProfile.Data.UserProfile.Name,
        UpdatedAt:      tempResponse.DcddCreateUserProfile.Data.UserProfile.UpdatedAt,
        UserId:         tempResponse.DcddCreateUserProfile.Data.UserProfile.UserId,
        GradeId:        tempResponse.DcddCreateUserProfile.Data.UserProfile.GradeId,
        SchoolId:       tempResponse.DcddCreateUserProfile.Data.UserProfile.SchoolId,
        ProfilePicture: tempResponse.DcddCreateUserProfile.Data.UserProfile.ProfilePicture,
        CreatedAt:      tempResponse.DcddCreateUserProfile.Data.UserProfile.CreatedAt,
        Dob:           tempResponse.DcddCreateUserProfile.Data.UserProfile.Dob,
        Dzongkhag: &model.Dzongkhag{
            ID:   tempResponse.DcddCreateUserProfile.Data.UserProfile.Dzongkhag.ID,
            Name: tempResponse.DcddCreateUserProfile.Data.UserProfile.Dzongkhag.Name,
        },
        Eccd: &model.Eccd{
            ID:        tempResponse.DcddCreateUserProfile.Data.UserProfile.Eccd.ID,
            Name:      tempResponse.DcddCreateUserProfile.Data.UserProfile.Eccd.Name,
            Sort:      tempResponse.DcddCreateUserProfile.Data.UserProfile.Eccd.Sort,     
        },
        Grade: &model.Grade{
            ID:   tempResponse.DcddCreateUserProfile.Data.UserProfile.Grade.ID,
            Name: tempResponse.DcddCreateUserProfile.Data.UserProfile.Grade.Name,            
        },  
        School: &model.School{
            ID:        tempResponse.DcddCreateUserProfile.Data.UserProfile.School.ID,
            PvtPublic: tempResponse.DcddCreateUserProfile.Data.UserProfile.School.PvtPublic,
            Name:      tempResponse.DcddCreateUserProfile.Data.UserProfile.School.Name,  
        },     
    }
    fmt.Println("formatted profile:", formattedProfile)
    return formattedProfile, nil
}

func GetUserProfile(userID uuid.UUID) (*model.UserProfile, error) {
 fmt.Println("Fetching profile for user ID:", userID)
	userServiceClient := graphql.NewClient(config.UserServiceApiUrl())
	req := graphql.NewRequest(`
		query DcddFetchProfileByUserId($user_id: UUID!) {
            dcddFetchProfileByUserId(user_id: $user_id) {
                data {
                    user_profile {
                        cid
                        created_at
                        dob
                        dzongkhag_id
                        eccd_id
                        gender
                        grade_id
                        id
                        name
                        profile_picture
                        school_id
                        updated_at
                        user_id
                        dzongkhag {
                            id
                            name
                        }
                        eccd {
                            id
                            name
                            sort
                        }
                        grade {
                            id
                            name
                        }
                        school {
                            id
                            pvt_public
                            name
                        }
                    }
                }
                error {
                    code
                    field
                    message
                }
            }
        }
    `)

	// Set the variable
	req.Var("user_id", userID)
	req.Header.Set("Cache-Control", "no-cache")

	// Define response struct
	var tempResponse struct {
        DcddFetchProfileByUserId struct {
            Data struct {
                GetUserProfile struct {
                    DzongkhagId uuid.UUID   `json:"dzongkhag_id"` 
                    ID         uuid.UUID `json:"id"`
                    Gender    string   `json:"gender"`
                    EccdId    uuid.UUID   `json:"eccd_id"`
                    Cid       string   `json:"cid"`
                    Name      string   `json:"name"`
                    UpdatedAt time.Time   `json:"updated_at"`
                    UserId    uuid.UUID   `json:"user_id"`
                    GradeId   uuid.UUID   `json:"grade_id"`
                    SchoolId  uuid.UUID   `json:"school_id"`
                    ProfilePicture string   `json:"profile_picture"`
                    CreatedAt time.Time   `json:"created_at"`
                    Dob     *time.Time   `json:"dob"`
                    Dzongkhag struct {
                        ID   uuid.UUID `json:"id"`
                        Name string `json:"name"`
                    } `json:"dzongkhag"`
                    Eccd struct {
                        ID        uuid.UUID `json:"id"`
                        Name      string `json:"name"`
                        Sort      string `json:"sort"`
                    } `json:"eccd"`
                    Grade struct {
                        ID   uuid.UUID `json:"id"`
                        Name string `json:"name"`
                    } `json:"grade"`
                    School struct {
                        ID        uuid.UUID `json:"id"`
                        PvtPublic string `json:"pvt_public"`
                        Name      string `json:"name"`  
                    } `json:"school"`   
                } `json:"user_profile"`
            } `json:"data"`
            Error struct {
                Code    string `json:"code"`
                Field   string `json:"field"`
                Message string `json:"message"`
            } `json:"error"`
        } `json:"dcddFetchProfileByUserId"`
    }
	// Execute the request
	err := userServiceClient.Run(context.Background(), req, &tempResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user profile : %v", err)
	}
	// Check for errors in response
	if tempResponse.DcddFetchProfileByUserId.Error.Message != "" {
		return nil, fmt.Errorf("error from video service: %s", tempResponse.DcddFetchProfileByUserId.Error.Message)
	}

	formattedProfile := &model.UserProfile{
        DzongkhagId:    tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.DzongkhagId,
        ID:             tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.ID,
        Gender:         tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Gender,
        EccdId:         tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.EccdId,
        Cid:            tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Cid,
        Name:           tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Name,
        UpdatedAt:      tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.UpdatedAt,
        UserId:         tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.UserId,
        GradeId:        tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.GradeId,
        SchoolId:       tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.SchoolId,
        ProfilePicture: tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.ProfilePicture,
        CreatedAt:      tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.CreatedAt,
        Dob:           tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Dob,
        Dzongkhag: &model.Dzongkhag{
            ID:   tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Dzongkhag.ID,
            Name: tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Dzongkhag.Name,
        },
        Eccd: &model.Eccd{
            ID:        tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Eccd.ID,
            Name:      tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Eccd.Name,
            Sort:      tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Eccd.Sort,     
        },
        Grade: &model.Grade{
            ID:   tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Grade.ID,
            Name: tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.Grade.Name,            
        },  
        School: &model.School{
            ID:        tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.School.ID,
            PvtPublic: tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.School.PvtPublic,
            Name:      tempResponse.DcddFetchProfileByUserId.Data.GetUserProfile.School.Name,  
        },     
    }

    fmt.Println("formatted profile:", formattedProfile)

    return formattedProfile, nil
}


