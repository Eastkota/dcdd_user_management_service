package model

import (
    "time"

    "github.com/google/uuid"
)

type DcddUser struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name           string    `gorm:"type:varchar(100);not null" json:"name"`
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
    ID                      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name                    string    `gorm:"type:varchar(100)" json:"name"`
    ProfilePicture          string    `gorm:"type:varchar(255)" json:"profile_picture"`
    Gender                  string    `gorm:"type:char(1)" json:"gender"`
    UserId                  uuid.UUID `gorm:"type:uuid" json:"user_id"`
    Cid                     string    `gorm:"type:varchar(50)" json:"cid"`
    SchoolId                uuid.NullUUID    `gorm:"type:varchar(32)" json:"school_id"`
	GradeId                 uuid.NullUUID    `gorm:"type:varchar(32)" json:"grade_id"`
	EccdId                  uuid.NullUUID    `gorm:"type:varchar(32)" json:"eccd_id"`
    DzongkhagId             uuid.UUID    `gorm:"type:varchar(32)" json:"dzongkhag_id"`
    Dob                     *time.Time   `gorm:"type:date" json:"dob"`
    CreatedAt               time.Time `json:"created_at"`
    UpdatedAt               time.Time `json:"updated_at"`
    School                  *School    `gorm:"foreignKey:SchoolId;references:ID" json:"school"`
	Grade                   *Grade     `gorm:"foreignKey:GradeId;references:ID" json:"grade"`
	Eccd      *Eccd      `gorm:"foreignKey:EccdId;references:ID" json:"eccd"`
	Dzongkhag *Dzongkhag `gorm:"foreignKey:DzongkhagId;references:ID" json:"dzongkhag"`
	DcddUser *DcddUser `gorm:"foreignKey:UserId;references:ID" json:"user"`
}

func (UserProfile) TableName() string {
    return "dcdd_user_data.dcdd_user_profiles"
}

type UserVideoPlaylist struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(100)" json:"name"`
    Favorites bool      `gorm:"type:boolean;default:false" json:"favorites"`
    ProfileId uuid.UUID `gorm:"type:uuid" json:"profile_id"`
    Videos []PlaylistVideo `gorm:"foreignKey:PlaylistId;references:ID" json:"videos"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (UserVideoPlaylist) TableName() string {
    return "dcdd_user_data.dcdd_video_playlists"
}