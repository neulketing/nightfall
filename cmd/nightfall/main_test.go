package main

import (
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	// Version should be set
	if version == "" {
		t.Error("version not set")
	}
}

func TestInitWorkspace(t *testing.T) {
	workspacePath := "/tmp/nightfall-test-init"
	initWorkspace(workspacePath)

	// Verify workspace directory exists
	entries, err := os.ReadDir(workspacePath)
	if err != nil {
		t.Errorf("failed to read workspace: %v", err)
	}

	// Verify subdirectories exist
	expectedDirs := []string{"logs", "config", "data", "targets", "exploits", "payloads", "c2", "reports"}
	for _, dir := range expectedDirs {
		found := false
		for _, entry := range entries {
			if entry.Name() == dir {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected directory %s not found", dir)
		}
	}
}
