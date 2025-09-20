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
		mutation DcddCreateDcddUserProfile($input: DcddUserProfileInput) {
            dcddCreateDcddUserProfile(input: $input) {
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
        DcddCreateDcddUserProfile struct {
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
        } `json:"dcddCreateDcddUserProfile"`
    }

	// Execute the request
	err := userServiceClient.Run(context.Background(), req, &tempResponse)
    fmt.Println("Response from user service:", err, tempResponse)
	if err != nil {
		return nil, err
	}
	// Check for errors in response
	if tempResponse.DcddCreateDcddUserProfile.Error.Message != "" {
		return nil, fmt.Errorf("error from video service: %s", tempResponse.DcddCreateDcddUserProfile.Error.Message)
	}


	formattedProfile := &model.UserProfile{
        DzongkhagId:    tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.DzongkhagId,
        ID:             tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.ID,
        Gender:         tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Gender,
        EccdId:         tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.EccdId,
        Cid:            tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Cid,
        Name:           tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Name,
        UpdatedAt:      tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.UpdatedAt,
        UserId:         tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.UserId,
        GradeId:        tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.GradeId,
        SchoolId:       tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.SchoolId,
        ProfilePicture: tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.ProfilePicture,
        CreatedAt:      tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.CreatedAt,
        Dob:           tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Dob,
        Dzongkhag: &model.Dzongkhag{
            ID:   tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Dzongkhag.ID,
            Name: tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Dzongkhag.Name,
        },
        Eccd: &model.Eccd{
            ID:        tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Eccd.ID,
            Name:      tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Eccd.Name,
            Sort:      tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Eccd.Sort,     
        },
        Grade: &model.Grade{
            ID:   tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Grade.ID,
            Name: tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.Grade.Name,            
        },  
        School: &model.School{
            ID:        tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.School.ID,
            PvtPublic: tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.School.PvtPublic,
            Name:      tempResponse.DcddCreateDcddUserProfile.Data.UserProfile.School.Name,  
        },     
    }
    fmt.Println("formatted profile:", formattedProfile)
    return formattedProfile, nil
}

func GetUserProfile(userID uuid.UUID) (*model.UserProfile, error) {
 fmt.Println("Fetching profile for user ID:", userID)
	userServiceClient := graphql.NewClient(config.UserServiceApiUrl())
	req := graphql.NewRequest(`
		query DcddFetchProfileByDcddUserId($user_id: UUID!) {
            dcddFetchProfileByDcddUserId(user_id: $user_id) {
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
        DcddFetchProfileByDcddUserId struct {
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
        } `json:"dcddFetchProfileByDcddUserId"`
    }
	// Execute the request
	err := userServiceClient.Run(context.Background(), req, &tempResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user profile : %v", err)
	}
	// Check for errors in response
	if tempResponse.DcddFetchProfileByDcddUserId.Error.Message != "" {
		return nil, fmt.Errorf("error from video service: %s", tempResponse.DcddFetchProfileByDcddUserId.Error.Message)
	}

	formattedProfile := &model.UserProfile{
        DzongkhagId:    tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.DzongkhagId,
        ID:             tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.ID,
        Gender:         tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Gender,
        EccdId:         tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.EccdId,
        Cid:            tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Cid,
        Name:           tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Name,
        UpdatedAt:      tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.UpdatedAt,
        UserId:         tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.UserId,
        GradeId:        tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.GradeId,
        SchoolId:       tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.SchoolId,
        ProfilePicture: tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.ProfilePicture,
        CreatedAt:      tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.CreatedAt,
        Dob:           tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Dob,
        Dzongkhag: &model.Dzongkhag{
            ID:   tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Dzongkhag.ID,
            Name: tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Dzongkhag.Name,
        },
        Eccd: &model.Eccd{
            ID:        tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Eccd.ID,
            Name:      tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Eccd.Name,
            Sort:      tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Eccd.Sort,     
        },
        Grade: &model.Grade{
            ID:   tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Grade.ID,
            Name: tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.Grade.Name,            
        },  
        School: &model.School{
            ID:        tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.School.ID,
            PvtPublic: tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.School.PvtPublic,
            Name:      tempResponse.DcddFetchProfileByDcddUserId.Data.GetUserProfile.School.Name,  
        },     
    }

    fmt.Println("formatted profile:", formattedProfile)

    return formattedProfile, nil
}


