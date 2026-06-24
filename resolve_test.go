package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestSamples runs the program as a subprocess on each sample file and verifies the output.
func TestSamples(t *testing.T) {
	// First, build the binary so we can execute it reliably
	binaryPath := filepath.Join(t.TempDir(), "tetris-optimizer")
	buildCmd := exec.Command("go", "build", "-o", binaryPath, ".")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	tests := []struct {
		name         string
		filePath     string
		shouldError  bool
		expectedSize int // Expected square board dimensions if successful, 0 if error
	}{
		{"Good 00 (1 tetromino)", "samples/good00.txt", false, 2},
		{"Good 01 (4 tetrominoes)", "samples/good01.txt", false, 5},
		{"Good 02 (6 tetrominoes)", "samples/good02.txt", false, 6},
		{"Good 03 (11 tetrominoes)", "samples/good03.txt", false, 7},
		{"Hard (12 tetrominoes)", "samples/hard.txt", false, 7},
		{"Bad 00 (empty or too short)", "samples/bad00.txt", true, 0},
		{"Bad 01 (invalid structure)", "samples/bad01.txt", true, 0},
		{"Bad 02 (disconnected tetromino)", "samples/bad02.txt", true, 0},
		{"Bad 03 (too many blocks)", "samples/bad03.txt", true, 0},
		{"Bad 04 (five blocks in one)", "samples/bad04.txt", true, 0},
		{"Bad Format (random invalid chars)", "samples/badformat.txt", true, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Resolve absolute path to the sample file
			absPath, err := filepath.Abs(tc.filePath)
			if err != nil {
				t.Fatalf("Failed to get absolute path: %v", err)
			}

			// Run the compiled binary on the sample file
			cmd := exec.Command(binaryPath, absPath)
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err = cmd.Run()
			if err != nil && !tc.shouldError {
				t.Fatalf("Command failed unexpectedly: %v\nStderr: %s", err, stderr.String())
			}

			outStr := strings.TrimSpace(stdout.String())

			if tc.shouldError {
				if outStr != "ERROR" {
					t.Errorf("Expected 'ERROR', got: %q", outStr)
				}
			} else {
				if outStr == "ERROR" {
					t.Fatalf("Expected valid layout, got ERROR")
				}
				lines := strings.Split(outStr, "\n")
				// Validate dimensions (must be a square of tc.expectedSize)
				if len(lines) != tc.expectedSize {
					t.Errorf("Expected height of %d, got %d. Output:\n%s", tc.expectedSize, len(lines), outStr)
				}
				for _, line := range lines {
					if len(line) != tc.expectedSize {
						t.Errorf("Expected width of %d, got %d in line: %q", tc.expectedSize, len(line), line)
					}
				}
			}
		})
	}
}
