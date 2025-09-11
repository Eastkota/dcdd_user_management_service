package helpers

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

)

// GenerateLoginLoginId generates a unique dynamic-length login ID
func GenerateLoginId(length int) string {
	if length <= 0 || length < 6 {
		return ""
	}
	// Create a new random generator with a custom seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate the first digit (non-zero)
	firstDigit := strconv.Itoa(r.Intn(9) + 1) // Ensures the range is 1-9

	// Generate the remaining digits to fit the length (6 - first digit)
	var buffer bytes.Buffer
	for i := 1; i < length; i++ {
		buffer.WriteString(strconv.Itoa(r.Intn(10)))
	}
	// Combine the first digit with the remaining digits
	return firstDigit + buffer.String()[:length-1]
}
