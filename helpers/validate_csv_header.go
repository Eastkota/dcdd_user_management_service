package helpers

import (
    "time"
    "fmt"
    "strings"

    "github.com/google/uuid"
)

// Parses a string into a uuid.NullUUID.
func ParseUUIDStringToNullUUID(s string) (uuid.NullUUID, error) {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return uuid.NullUUID{}, nil
    }
    
    parsedUUID, err := uuid.Parse(trimmed)
    if err != nil {
        return uuid.NullUUID{}, fmt.Errorf("invalid UUID format: %s", s)
    }
    return uuid.NullUUID{UUID: parsedUUID, Valid: true}, nil
}

// Parses a string into a *time.Time.
func ParseDateString(s string) (*time.Time, error) {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return nil, nil
    }
    
    t, err := time.Parse("2006-01-02", trimmed)
    if err != nil {
        return nil, fmt.Errorf("invalid date format: %s", s)
    }
    return &t, nil
}

func ValidateCSVHeader(header []string) bool {
    expected := []string{"Name", "MobileNo", "Email", "Gender", "Password","StudentId","Category","SchoolId","GradeId","EccdId","Dob","DzongkhagId","Cid"}
    if len(header) != len(expected) {
        return false
    }
    for i, h := range header {
        if h != expected[i] {
            return false
        }
    }
    return true
}