package string

import (
	"fmt"

	"github.com/google/uuid"
)

// Generate username with 8 digit uuid
// as appended suffix
func GenerateRandomUsername() string {
	uuid := uuid.New().String()
	return fmt.Sprintf("user_%v", uuid[0:8])
}
