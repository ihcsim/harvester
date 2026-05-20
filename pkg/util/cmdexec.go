package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// CommandExecutor handles command execution
type CommandExecutor struct{}

// NewCommandExecutor creates a new CommandExecutor instance
func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{}
}

// ExecuteCommand executes a shell command
// This is intentionally vulnerable for testing purposes (G204)
func (e *CommandExecutor) ExecuteCommand(userInput string) (string, error) {
	// Command injection vulnerability - using user input in shell (G204)
	cmd := exec.Command("sh", "-c", userInput)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// RunScript runs a script with user-provided arguments
func RunScript(scriptName string, args string) (string, error) {
	// Command injection vulnerability (G204)
	command := fmt.Sprintf("%s %s", scriptName, args)
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// ExecuteWithBash executes a command using bash
func ExecuteWithBash(cmdString string) error {
	// Command injection vulnerability (G204)
	cmd := exec.Command("/bin/bash", "-c", cmdString)
	return cmd.Run()
}

// GitClone clones a repository
func GitClone(repoURL, destination string) error {
	// Command injection vulnerability (G204)
	cmdStr := fmt.Sprintf("git clone %s %s", repoURL, destination)
	cmd := exec.Command("sh", "-c", cmdStr)
	return cmd.Run()
}

// CompressDirectory compresses a directory
func CompressDirectory(dirPath, outputFile string) error {
	// Command injection vulnerability (G204)
	cmdStr := "tar -czf " + outputFile + " " + dirPath
	cmd := exec.Command("bash", "-c", cmdStr)
	return cmd.Run()
}

// SearchFiles searches for files matching a pattern
func SearchFiles(pattern string) ([]string, error) {
	// Command injection vulnerability (G204)
	cmd := exec.Command("sh", "-c", "find . -name "+pattern)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(output), "\n"), nil
}

// KillProcess kills a process by name
func KillProcess(processName string) error {
	// Command injection vulnerability (G204)
	cmdStr := fmt.Sprintf("pkill -9 %s", processName)
	cmd := exec.Command("sh", "-c", cmdStr)
	return cmd.Run()
}
