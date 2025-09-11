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
    Email          string    `gorm:"type:varchar(100);unique" json:"email"`
    MobileNo       string    `gorm:"type:varchar(20);unique" json:"mobile_no"`
    Password       string    `gorm:"type:text;not null" json:"password_hash"`
    Status         string    `gorm:"type:varchar(50);default:active;not null" json:"status"`
    Category  	   string    `gorm:"not null" json:"category"`
    StudentId 	   string    `gorm:"type:varchar(50);unique" json:"student_id"`
    LoginId        string    `gorm:"type:varchar(20);unique;not null" json:"login_id"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

func (DcddUser) TableName() string {
    return "dcdd_auth.dcdd_users"
}

type UserResult struct {
    DcddUser *DcddUser `json:"user"`
}

type UserProfile struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255)" json:"name"`
	ProfilePicture string    `gorm:"profile_picture" json:"profile_picture,omitempty"`
	Gender         string    `gorm:"type:varchar(20)" json:"gender,omitempty"`
	UserId         uuid.UUID `gorm:"type:uuid" json:"user_id"`
	SchoolId       uuid.UUID     `gorm:"type:uuid" json:"school_id,omitempty"`
	GradeId        uuid.UUID     `gorm:"type:uuid" json:"grade_id,omitempty"`
	EccdId         uuid.UUID     `gorm:"type:uuid" json:"eccd_id,omitempty"`
	Dob            *time.Time    `gorm:"type:dob" json:"dob,omitempty"`
	DzongkhagId    uuid.UUID     `gorm:"type:uuid" json:"dzongkhag_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	Cid            string    `gorm:"type:varchar(50)" json:"cid,omitempty"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	School    *School    `gorm:"foreignKey:SchoolId;references:ID" json:"school,omitempty"`
	Grade     *Grade     `gorm:"foreignKey:GradeId;references:ID" json:"grade,omitempty"`
	Eccd      *Eccd      `gorm:"foreignKey:EccdId;references:ID" json:"eccd,omitempty"`
	Dzongkhag *Dzongkhag `gorm:"foreignKey:DzongkhagId;references:ID" json:"dzongkhag"`
}

func (UserProfile) TableName() string {
    return "dcdd_user_data.dcdd_user_profiles"
}
type School struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	PvtPublic string `gorm:"type:varchar(50)" json:"pvt_public"`
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
	Sort      int64  `gorm:"type:varchar(50)" json:"sort"`
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


