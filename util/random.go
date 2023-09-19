package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random username
func RandomUsername() string {
	return fmt.Sprintf("%v%v", RandomString(6), RandomInt(999, 10000))
}

// Generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%v@email.com", RandomString(8))
}

// Generate a random role name
func RandomRoleID() int64 {
	numbers := []int64{1, 2}
	return numbers[rand.Intn(2)]
}

// Generates a random tag
func RandomCategory() string {
	category := []string{
		"golang", "python", "javascript",
		"java", "kotlin", "swift",
		"csharp", "php", "ruby",
		"rust", "typescript", "html",
		"css", "sql", "docker",
	}
	return category[rand.Intn(15)]
}
