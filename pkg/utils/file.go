package utils

import (
	"os"
)

// IsFileExists return the file is or not exists
func IsFileExists(filePath string) bool {
    _, err := os.Stat(filePath)

    return os.IsNotExist(err) == false
}
