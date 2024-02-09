package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomID(t *testing.T) {
	t.Run("Default type - alphanumeric", func(t *testing.T) {
		// Generate a random ID with default type (alphanumeric)
		id := RandomID(10)

		// Verify length
		require.Equal(t, 10, len(id))

		// Verify character type
		for _, char := range id {
			isAlphaNum := 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || '0' <= char && char <= '9'
			require.True(t, isAlphaNum)
		}
	})

	t.Run("Custom type - numeric", func(t *testing.T) {
		// Generate a random ID with numeric type
		id := RandomID(8, SeedTypeNumber)

		// Verify length
		require.Equal(t, 8, len(id))

		// Verify character type
		for _, char := range id {
			isNumeric := '0' <= char && char <= '9'
			require.True(t, isNumeric)
		}
	})

	t.Run("Custom type - alpha", func(t *testing.T) {
		// Generate a random ID with alpha type
		id := RandomID(6, SeedTypeAlpha)

		// Verify length
		require.Equal(t, 6, len(id))

		// Verify character type
		for _, char := range id {
			isAlpha := ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
			require.True(t, isAlpha)
		}
	})

	t.Run("Custom type - alphanum without similar characters", func(t *testing.T) {
		// Generate a random ID with alphanumnosim type
		id := RandomID(15, SeedTypeAlphaNumNoSimilarity)

		// Verify length
		require.Equal(t, 15, len(id))

		// Verify character type
		for _, char := range id {
			isAlphaNumNoSimilar := ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('2' <= char && char <= '9')
			require.True(t, isAlphaNumNoSimilar)
		}
	})
}

func TestPasswordValidate(t *testing.T) {
	// Define valid and invalid passwords with different edge cases
	validPasswords := map[string]bool{
		"Passw0rd1!":   true, // upper, lower, symbol, number
		"passw0rd!":    false,
		"123456789@":   false,
		"PASSWORD123!": false,
		"Aa1!1224":     true,  // minimum length
		"password":     false, // no uppercase
		"PASSWORD":     false, // no lowercase
		"1234567890":   false, // no symbol
		"qwertyuiop":   false, // no number
		"!@#$%^&*()":   false, // no letter
		"short":        false, // below minimum length
		"":             false, // empty password
	}

	// Test valid passwords
	for password, expected := range validPasswords {
		t.Run(password, func(t *testing.T) {
			isValid := PasswordValidate(password, 8)
			require.Equal(t, expected, isValid)
		})
	}
}
