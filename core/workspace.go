package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Workspace represents the Nightfall workspace
type Workspace struct {
	Path      string    `json:"path"`
	Created   time.Time `json:"created"`
	Targets   []string  `json:"targets"`
	Campaigns []string  `json:"campaigns"`
}

// LoadWorkspace loads the workspace configuration
func LoadWorkspace(path string) (*Workspace, error) {
	configPath := filepath.Join(path, "config", "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var ws Workspace
	if err := json.Unmarshal(data, &ws); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &ws, nil
}

// SaveWorkspace saves the workspace configuration
func SaveWorkspace(ws *Workspace) error {
	data, err := json.MarshalIndent(ws, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	configPath := filepath.Join(ws.Path, "config", "config.json")
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}
