package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: vcl_formatter <file.vcl>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file %s: %v\n", os.Args[1], err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	indentLevel := 0
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "if ") || strings.HasPrefix(line, "sub vcl_") {
			if strings.HasSuffix(line, "{") {
				indentLevel++
			}
		} else if strings.HasPrefix(line, "}") {
			if indentLevel > 0 {
				indentLevel--
			}
		}

		indentedLine := fmt.Sprintf("%s%s\n", strings.Repeat("    ", indentLevel), line)
		writer.WriteString(indentedLine)

		if strings.HasSuffix(line, "{") {
			indentLevel++
		}
	}
}
