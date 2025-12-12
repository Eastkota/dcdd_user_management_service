package model

import (
    "time"
    "github.com/google/uuid"
)

type ContextKey string

const (
    UserKey    ContextKey = "user"
    RequestKey ContextKey = "http_request"
)

type DcddUser struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    UserIdentifier string    `gorm:"type:varchar(32);unique;not null" json:"user_identifier"`
    Email          *string    `gorm:"type:varchar(100);unique" json:"email"`
    MobileNo       *string    `gorm:"type:varchar(20);unique" json:"mobile_no"`
    Password       string    `gorm:"type:text;not null" json:"password_hash"`
    Status         string    `gorm:"type:varchar(50);default:active;not null" json:"status"`
    Category  	   string    `gorm:"not null" json:"category"`
    StudentId 	   string    `gorm:"type:varchar(50);unique" json:"student_id"`
    LoginId        string    `gorm:"type:varchar(20);unique;not null" json:"login_id"`
	CreatedAt      time.Time `gorm:"type:timestamptz" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamptz" json:"updated_at"`

	UserProfile	   *UserProfile `gorm:"foreignKey:UserId;references:ID" json:"user_profile"`
}

func (DcddUser) TableName() string {
    return "dcdd_auth.dcdd_users" 
}

type UserProfile struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255)" json:"name"`
	ProfilePicture string    `gorm:"profile_picture" json:"profile_picture,omitempty"`
	Gender         string    `gorm:"type:varchar(20)" json:"gender,omitempty"`
	UserId         uuid.UUID `gorm:"type:uuid" json:"user_id"`
	SchoolId       uuid.UUID     `gorm:"type:uuid;default:null" json:"school_id,omitempty"`
	GradeId        uuid.UUID     `gorm:"type:uuid;default:null" json:"grade_id,omitempty"`
	EccdId         uuid.UUID     `gorm:"type:uuid;default:null" json:"eccd_id,omitempty"`
	Dob            *time.Time    `gorm:"type:date" json:"dob,omitempty"`
	DzongkhagId    uuid.UUID     `gorm:"type:uuid" json:"dzongkhag_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	Cid            string    `gorm:"type:varchar(50)" json:"cid,omitempty"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	School    *School    `gorm:"foreignKey:SchoolId;references:ID" json:"school,omitempty"`
	Grade     *Grade     `gorm:"foreignKey:GradeId;references:ID" json:"grade,omitempty"`
	Eccd      *Eccd      `gorm:"foreignKey:EccdId;references:ID" json:"eccd,omitempty"`
	Dzongkhag *Dzongkhag `gorm:"foreignKey:DzongkhagId;references:ID" json:"dzongkhag"`

	User *DcddUser `gorm:"foreignKey:UserId;references:ID" json:"user"`
}

func (UserProfile) TableName() string {
    return "dcdd_user_data.dcdd_user_profiles"
}
type DcddUserAndProfile struct {
	User *DcddUser `json:"user"`
	UserProfile *UserProfile `json:"user_profile"`
	 
}
type School struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	PvtPublic string `gorm:"type:varchar(50)" json:"pvt_public"`
	DzonghkhagId uuid.UUID `gorm:"type:uuid" json:"dzongkhag_id"`
}
type SchoolResult struct {
	School []School `gorm:"type:uuid;primaryKey" json:"school"`
}
func (School) TableName() string {
	return "dcdd_user_data.schools"
}
type Grade struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}
type GradeResult struct {
	Grades []Grade `gorm:"type:uuid;primaryKey" json:"grade"`
}
func (Grade) TableName() string {
	return "dcdd_user_data.grades"
}

type Eccd struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Sort      string  `gorm:"type:varchar(50)" json:"sort"`
	DzonghkhagId uuid.UUID `gorm:"type:uuid" json:"dzongkhag_id"`
}
type EccdResult struct {
	Eccd []Eccd `gorm:"type:uuid;primaryKey" json:"eccd"`
}
func (Eccd) TableName() string {
	return "dcdd_user_data.eccds"
}
type Dzongkhag struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}

type DzongkhagResult struct {
	Dzongkhags []Dzongkhag `gorm:"type:uuid;primaryKey" json:"dzongkhags"`
}

func (Dzongkhag) TableName() string {
	return "dcdd_user_data.dzongkhags"
}

type DcddUserPagination struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	Limit       int `json:"limit"`
}

type FetchAllDcddUsersResult struct {
    Users     []DcddUser `json:"user"`
    Pagination *DcddUserPagination   `json:"pagination"`
}

type UserActivity struct {
    ID  uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    Activity string `gorm:"type:varchar" json:"activity"`
    UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Count   int     `json:"count" gorm:"type:integer"`
    Month           int         `json:"month" gorm:"type:integer"`             
    Year            int         `json:"year" gorm:"type:integer"`

    User    *DcddUser   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (UserActivity) TableName() string {
    return "dcdd_auth.dcdd_user_activities"
}

type ActivityCounts map[string]int

type GroupedUserActivity struct {
    UserID          uuid.UUID       `json:"user_id"`
    Month           int             `json:"month"`
    Year            int             `json:"year"`
    ActivityCounts  ActivityCounts  `json:"activity_counts"` // Contains {"video watched": 10, "others": 12}
    TotalCount      int             `json:"total_count"`     // Sum of all counts (e.g., 22)
    User            *DcddUser `json:"user"`            // The preloaded user details
}


