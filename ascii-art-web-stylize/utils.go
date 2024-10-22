package main

import "os"

func ensureFiles(importantFiles map[int]string) bool {
	for _, fileName := range importantFiles {
		if _, err := os.Stat(fileName); err == nil {
			continue
		} else if err == os.ErrNotExist {
			return false
		}
		return false
	}
	return true
}
