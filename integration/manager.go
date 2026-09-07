package integration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IntegrationManager manages external tool integrations
type IntegrationManager struct {
	BaseDir string
	Tools   map[string]string
}

// NewIntegrationManager creates a new integration manager
func NewIntegrationManager(baseDir string) *IntegrationManager {
	return &IntegrationManager{
		BaseDir: baseDir,
		Tools: map[string]string{
			"atomic":    filepath.Join(baseDir, "integrations", "atomic-red-team"),
			"caldera":   filepath.Join(baseDir, "integrations", "caldera"),
			"mordor":    filepath.Join(baseDir, "integrations", "mordor"),
		},
	}
}

// RunAtomicTest executes an Atomic Red Team test
func (m *IntegrationManager) RunAtomicTest(technique string, target string) error {
	_, exists := m.Tools["atomic"]
	if !exists {
		return fmt.Errorf("atomic-red-team not found")
	}

	fmt.Printf("Running Atomic Red Team test: %s\n", technique)

	// In a real implementation, this would:
	// - Invoke Invoke-AtomicRedTeam PowerShell module
	// - Or use the Python runner
	// - Pass the technique ID and target
	// - Parse and return results

	return nil
}

// StartCaldera starts the CALDERA server
func (m *IntegrationManager) StartCaldera() error {
	_, exists := m.Tools["caldera"]
	if !exists {
		return fmt.Errorf("caldera not found")
	}

	fmt.Printf("Starting CALDERA server...\n")

	// In a real implementation, this would:
	// - Install CALDERA dependencies
	// - Configure the server
	// - Start the HTTP server
	// - Return the server URL

	return nil
}

// RunMordorScenario executes a Mordor scenario
func (m *IntegrationManager) RunMordorScenario(scenario string) error {
	_, exists := m.Tools["mordor"]
	if !exists {
		return fmt.Errorf("mordor not found")
	}

	fmt.Printf("Running Mordor scenario: %s\n", scenario)

	// In a real implementation, this would:
	// - Execute the scenario JSON file
	// - Parse the telemetry output
	// - Return results

	return nil
}

// CheckIntegration verifies an integration is available
func (m *IntegrationManager) CheckIntegration(tool string) (bool, string) {
	toolPath, exists := m.Tools[tool]
	if !exists {
		return false, "not configured"
	}

	if _, err := os.Stat(toolPath); os.IsNotExist(err) {
		return false, "directory not found"
	}

	// Check for required files
	if strings.Contains(tool, "atomic") {
		testFiles, _ := os.ReadDir(filepath.Join(toolPath, "tests"))
		return len(testFiles) > 0, "tests found"
	}

	if strings.Contains(tool, "caldera") {
		if _, err := os.Stat(filepath.Join(toolPath, "app")); os.IsNotExist(err) {
			return false, "app directory not found"
		}
		return true, "app directory found"
	}

	if strings.Contains(tool, "mordor") {
		scenarios, _ := os.ReadDir(filepath.Join(toolPath, "scenarios"))
		return len(scenarios) > 0, "scenarios found"
	}

	return true, "directory exists"
}
