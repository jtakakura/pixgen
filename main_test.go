package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainCLI(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "pixgen_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a sample input JSON file
	inputData := InputData{
		"test_image": {
			{
				"........",
				"..GG....",
				"..GG....",
				"........",
				"....GG..",
				"....GG..",
				"........",
				"........",
			},
		},
	}
	inputFilePath := filepath.Join(tempDir, "input.json")
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer inputFile.Close()

	if err := json.NewEncoder(inputFile).Encode(inputData); err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}

	// Run the CLI
	cmd := exec.Command("go", "run", "main.go", "-input", inputFilePath, "-output", tempDir)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("CLI execution failed: %v, stderr: %s", err, stderr.String())
	}

	// Check if the output file exists
	outputFilePath := filepath.Join(tempDir, "test_image.png")
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Fatalf("Output file not found: %s", outputFilePath)
	}
}
