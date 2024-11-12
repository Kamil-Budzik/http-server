package helpers

import (
	"os"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	t.Run("ValidFile", func(t *testing.T) {
		filePath := "testfile.txt"
		content := "Hello, World!\nThis is a test file.\n"
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}
		defer os.Remove(filePath) // Clean up the test file

		result, err := ReadFileLines(filePath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != content {
			t.Errorf("expected %v, got %v", content, result)
		}
	})

	t.Run("NonExistentFile", func(t *testing.T) {
		_, err := ReadFileLines("nonexistentfile.txt")
		if err == nil {
			t.Errorf("expected error for non-existent file, got nil")
		}
	})

	t.Run("EmptyFile", func(t *testing.T) {
		filePath := "emptyfile.txt"
		err := os.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			t.Fatalf("failed to create empty test file: %v", err)
		}
		defer os.Remove(filePath)

		result, err := ReadFileLines(filePath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != "" {
			t.Errorf("expected empty string, got %v", result)
		}
	})
}
