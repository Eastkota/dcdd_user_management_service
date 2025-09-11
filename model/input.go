package model

import (
    "github.com/google/uuid"
    "time"
)

type SignupInput struct {
    Email          string `json:"email"`
	MobileNo       string `json:"mobile_no"`
	Name           string `json:"name"`
	GradeId        uuid.NullUUID `json:"grade_id"`
	Gender         string `json:"gender"`
	Password       string `json:"password"`
	StudentId      string `json:"student_id"`
	SchoolId       uuid.NullUUID `json:"school_id"`
	EccdId         uuid.NullUUID `json:"eccd_id"`
	Dob            *time.Time `json:"dob"`
	Cid            string `json:"cid"`
	DzongkhagId    uuid.UUID `json:"dzongkhag_id"`
	Category       string `json:"category"`
}

type UpdatePasswordInput struct {
    CurrentPassword string `json:"current_password"`
    Password        string `json:"password"`
    ConfirmPassword string `json:"confirm_password"`
    UserId          uuid.UUID `json:"user_id"`
}

type UpdateSingleAuthDataInput struct {
    Field    string `json:"field"`
    Value    string `json:"value"`
    Password string `json:"password"`
}

type ResetPasswordInput struct {
    UserId          uuid.UUID `json:"user_id"`
    Password        string    `json:"password"`
    ConfirmPassword string    `json:"confirm_password"`
}

type UserProfileInput struct {
    Name           string `json:"name"`
    Gender         string `json:"gender"`
    ProfilePicture string `json:"profile_picture"`
    UserId         uuid.UUID `json:"user_id"`
    GradeId       uuid.NullUUID `json:"grade_id"`
    SchoolId      uuid.NullUUID `json:"school_id"`
    EccdId        uuid.NullUUID `json:"eccd_id"`
    DzongkhagId   uuid.UUID `json:"dzongkhag_id"`
    Dob           *time.Time `json:"dob"`
    Cid           string `json:"cid"`
}