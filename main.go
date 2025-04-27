package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
)

type InputData map[string][][]string

var charToColor = map[rune]color.RGBA{
	'.': {0, 0, 0, 0},         // transparent
	'l': {0, 0, 0, 255},       // black
	'r': {255, 0, 0, 255},     // red
	'g': {0, 255, 0, 255},     // green
	'b': {0, 0, 255, 255},     // blue
	'y': {255, 255, 0, 255},   // yellow
	'p': {128, 0, 128, 255},   // purple
	'c': {0, 255, 255, 255},   // cyan
	'w': {255, 255, 255, 255}, // white
	'L': {85, 85, 85, 255},    // light_black
	'R': {255, 128, 128, 255}, // light_red
	'G': {128, 255, 128, 255}, // light_green
	'B': {128, 128, 255, 255}, // light_blue
	'Y': {255, 255, 128, 255}, // light_yellow
	'P': {255, 128, 255, 255}, // light_purple
	'C': {128, 255, 255, 255}, // light_cyan
	'W': {170, 170, 170, 255}, // light_white
}

func main() {
	// Define command-line arguments
	inputPath := flag.String("input", "", "Path to the input JSON file (required)")
	outputDir := flag.String("output", ".", "Path to the output directory (optional, defaults to current directory)")
	flag.Parse()

	// Check for help flag
	if flag.Arg(0) == "-h" || flag.Arg(0) == "--help" {
		fmt.Println("Usage: pixgen -input <path> -output <dir>")
		fmt.Println("\nOptions:")
		fmt.Println("  -input <path>   Path to the input JSON file (required)")
		fmt.Println("  -output <dir>   Path to the output directory (optional, defaults to current directory)")
		os.Exit(0)
	}

	// Check required arguments
	if *inputPath == "" {
		log.Fatal("Error: -input flag is required")
	}

	// Load the JSON file
	file, err := os.Open(*inputPath)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer file.Close()

	var inputData InputData
	if err := json.NewDecoder(file).Decode(&inputData); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Create the output directory
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	fmt.Println("Input JSON successfully loaded and output directory prepared.")

	// Generate images for each key in the input data
	for key, imageDefinitions := range inputData {
		outputPath := filepath.Join(*outputDir, key+".png")
		if err := generateImage(imageDefinitions, outputPath); err != nil {
			log.Fatalf("Error generating image for key '%s': %v", key, err)
		}
	}
}

func generateImage(imageDefinitions [][]string, outputPath string) error {
	if len(imageDefinitions) == 0 {
		return fmt.Errorf("No image definitions provided")
	}

	// Determine the dimension of the images
	dimension := len(imageDefinitions[0])
	canvasSize := int(math.Ceil(math.Sqrt(float64(len(imageDefinitions))))) * dimension
	canvas := image.NewRGBA(image.Rect(0, 0, canvasSize, canvasSize))

	// Fill the canvas with transparent background
	for y := 0; y < canvasSize; y++ {
		for x := 0; x < canvasSize; x++ {
			canvas.Set(x, y, color.RGBA{0, 0, 0, 0})
		}
	}

	// Calculate grid size outside the loop
	gridSize := int(math.Ceil(math.Sqrt(float64(len(imageDefinitions)))))

	// Draw each image onto the canvas
	for i, definition := range imageDefinitions {
		img := image.NewRGBA(image.Rect(0, 0, dimension, dimension))
		for y, row := range definition {
			for x, char := range row {
				col, ok := charToColor[char]
				if !ok {
					return fmt.Errorf("Invalid character '%c' in image definition", char)
				}
				img.Set(x, y, col)
			}
		}

		// Calculate position on the canvas
		gridX := (i % gridSize) * dimension
		gridY := (i / gridSize) * dimension
		rect := image.Rect(gridX, gridY, gridX+dimension, gridY+dimension)
		draw.Draw(canvas, rect, img, image.Point{}, draw.Src)
	}

	// Save the canvas as a PNG file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, canvas); err != nil {
		return fmt.Errorf("Failed to encode PNG: %v", err)
	}

	return nil
}
