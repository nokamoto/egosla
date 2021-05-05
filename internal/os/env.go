package os

import "os"

// Getenv retrieves the value of the environment variable named by the key.
// If the value is empty, it returns the or value.
func GetenvOr(key string, or string) string {
	s := os.Getenv(key)
	if len(s) == 0 {
		return or
	}
	return s
}
