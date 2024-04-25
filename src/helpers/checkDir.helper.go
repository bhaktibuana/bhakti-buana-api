package helpers

import (
	"os"
)

// CheckDir Helper
/*
 * @param dir string
 * @returns
 */
func CheckDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return
		}
	}
}
