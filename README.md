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



- This is a brief overview of the available functions. Refer to the package documentation, unit test for more details and examples.
