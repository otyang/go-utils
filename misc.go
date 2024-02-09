package utils

import (
	"crypto/rand"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	SeedTypeAlphaNum             = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SeedTypeAlphaNumNoSimilarity = "2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ"
	SeedTypeAlpha                = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SeedTypeNumber               = "0123456789"
)

// RandomID generates a random string of specified size using a given type of characters.
//
// - strSize: desired length of the generated random string.
// - randType (optional): type of characters to use for randomization.
//   - Supports "alpha", "alphanum", "number", "alphanumnosim" (alphanum without similar characters).
//   - Defaults to "alphanum" (letters and numbers) if no type is specified.
//
// Returns:
// - A random string of the specified length and character type.
//
// Example usage:
//
//	randomID := RandomID(10)  // Generate a 10-character alphanum string.
//	randomID := RandomID(10, "number")  // Generate a 10-character numeric string.
func RandomID(strSize int, randType ...string) string {
	// lets set default rand type
	dictionary := SeedTypeAlphaNum

	if len(randType) > 0 {
		seed := strings.TrimSpace(randType[0])
		switch seed {
		case SeedTypeAlpha, SeedTypeAlphaNum, SeedTypeNumber, SeedTypeAlphaNumNoSimilarity:
			dictionary = seed
		}
	}

	bytes := make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

// PasswordValidate checks if a password meets the following criteria:
//
// * Minimum length of `pwdLen` characters (configurable)
// * Contains at least one uppercase letter (A-Z)
// * Contains at least one lowercase letter (a-z)
// * Contains at least one number (0-9)
// * Contains at least one symbol (including punctuation and special characters)
//
// This function iterates through each character in the password and uses
// Unicode category checks to categorize them. If any category is missing
// or the total character count falls below the minimum length, validation fails.
// Otherwise, the password is considered valid.
func PasswordValidate(pass string, pwdLen int) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < uint8(pwdLen) {
		return false
	}
	return true
}

// HashPassword takes a plain text password and returns a hashed version using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compares a plaintext password with a stored password hash using the bcrypt algorithm.
// It returns true if the passwords match, false otherwise.
func ComparePasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ToPointer accepts a value of any type and returns a pointer to that
// value. This allows you to achieve similar functionality as the
// built-in "&" operator, but with increased type safety and flexibility.
//
// Unlike "&", which only works with variables, ToPointer can be used
// with expressions, constants, and function return values. This makes it
// a more versatile tool for working with pointers in various contexts.
func ToPointer[T any](v T) *T {
	return &v
}

const (
	// Default format for timestamps used in identifiers and filenames
	defaultTimeFormat = "20060102150405"
)

// FormattedTime returns the current time formatted according to the
// specified format, or the default one if no format is provided.
func FormattedTime(format string) string {
	if format == "" {
		format = defaultTimeFormat
	}
	return time.Now().UTC().Format(format)
}
