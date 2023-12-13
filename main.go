package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if a file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: vcl_formatter <file.vcl>")
		os.Exit(1)
	}

	// Open the file
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Format the VCL code
	formattedLines := formatVCL(lines)

	// Output the formatted code
	for _, line := range formattedLines {
		fmt.Println(line)
	}
}

func formatVCL(lines []string) []string {
	var formattedLines []string
	indentLevel := 0

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// Handle the opening brace for specific structures
		if strings.HasPrefix(line, "if ") || strings.HasPrefix(line, "sub vcl_") {
			// Check if the next line contains only an opening brace
			if i+1 < len(lines) && strings.TrimSpace(lines[i+1]) == "{" {
				line += " {"
				i++ // Skip the next line as it's been merged with the current line
			}
		} else if strings.HasPrefix(line, "}") {
			// Decrease indent level for closing braces
			if indentLevel > 0 {
				indentLevel--
			}
		}

		// Add indentation
		indentedLine := strings.Repeat("    ", indentLevel) + line
		formattedLines = append(formattedLines, indentedLine)

		// Increase indent level for opening braces
		if strings.HasSuffix(line, "{") {
			indentLevel++
		}
	}

	return formattedLines
}
