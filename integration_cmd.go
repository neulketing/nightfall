package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runIntegration(command string, args []string) {
	workspace := ""
	for i, arg := range args {
		if arg == "--workspace" && i+1 < len(args) {
			workspace = args[i+1]
		}
	}

	if workspace == "" {
		fmt.Println("Error: --workspace required")
		return
	}

	// Check for integration tools
	tools := []string{"atomic", "caldera", "mordor"}
	for _, tool := range tools {
		toolPath := filepath.Join(workspace, "integrations", fmt.Sprintf("%s-%s", tool, tool))
		if _, err := os.Stat(toolPath); os.IsNotExist(err) {
			fmt.Printf("Cloning %s...\n", tool)
			cmd := exec.Command("git", "clone", fmt.Sprintf("https://github.com/%s/%s.git", getToolOwner(tool), getToolRepo(tool)), toolPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("Failed to clone %s: %v\n", tool, err)
			}
		}
	}

	fmt.Println("Integration tools ready")
}

func getToolOwner(tool string) string {
	switch tool {
	case "atomic":
		return "redcanaryco"
	case "caldera":
		return "mitre"
	case "mordor":
		return "Cyb3rWard0g"
	default:
		return ""
	}
}

func getToolRepo(tool string) string {
	switch tool {
	case "atomic":
		return "atomic-red-team"
	case "caldera":
		return "caldera"
	case "mordor":
		return "mordor"
	default:
		return ""
	}
}
