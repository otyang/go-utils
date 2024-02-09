# go-utils

The `utils` package provides a collection of useful functions and constants for common tasks in Go applications. It includes tools for generating random strings, validating passwords, hashing passwords, working with pointers, and formatting timestamps.

**Installation**

```bash
go get github.com/otyang/go-utils
```

**Features**

- **`RandomID`:** Generates random strings of specified size and character type (alphanumeric, numeric, etc.)
- **`PasswordValidate`:** Checks if a password meets predefined complexity requirements.
- **`HashPassword`:** Creates a bcrypt hash from a plain text password.
- **`ComparePasswordAndHash`:** Compares a plain text password with a stored bcrypt hash.
- **`ToPointer`:** Safely converts a value to a pointer.
- **`FormattedTime`:** Formats the current time according to a specified or default format.

**Usage**

**Random Strings:**
```go
package main

import (
    "fmt"
    "github.com/your-username/utils"
)

func main() {
    randID := utils.RandomID(10) // Generate a 10-character alphanumeric string
    fmt.Println("Random ID:", randID)

    randIDNoSim := utils.RandomID(10, utils.SeedTypeAlphaNumNoSimilarity) // No similar characters
    fmt.Println("Random ID (no similar characters):", randIDNoSim)
}
```

**Password Validation:**

```go
func main() {
    valid := utils.PasswordValidate("StrongPassword!", 8) // Should be valid
    fmt.Println("Password is valid:", valid)

    valid = utils.PasswordValidate("weak123", 8) // Too short, lacks complexity
    fmt.Println("Password is valid:", valid)
}
```

**Password Hashing and Comparison:**
```go
func main() {
    password := "MySecurePassword"
    hash, err := utils.HashPassword(password)
    if err != nil {
        // Handle error
    }

    valid := utils.ComparePasswordAndHash(password, hash)
    fmt.Println("Password matches hash:", valid)
}
```

**Pointers:**
```go
func main() {
    value := 42
    ptr := utils.ToPointer(value) // Get a pointer to the value
    fmt.Println("Value:", *ptr)
}
```

**Timestamp Formatting:**
```go
func main() {
    formattedTime := utils.FormattedTime("") // Use default format
    fmt.Println("Formatted time:", formattedTime)

    customTime := utils.FormattedTime("YYYY-MM-DD") // Custom format
    fmt.Println("Custom formatted time:", customTime)
}
```

**Additional Notes**
- This is a brief overview of the available functions. Refer to the package documentation, unit test for more details and examples.
