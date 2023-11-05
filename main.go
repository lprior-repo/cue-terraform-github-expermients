package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
  "text/template"
	"strings"
)

// SLO represents the structure of our SLO data
type SLO struct {
	Name             string
	Description      string
	NumeratorQuery   string
	DenominatorQuery string
	SevenDayThresh   string
	ThirtyDayThresh  string
}

func main() {
	// Open the file
	file, err := os.Open("slo.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Read the content of the file
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the strings.Builder content to a string
	markdownContent := content.String()

	// Regular expression pattern to match the sections and their content
	pattern := regexp.MustCompile(`### ([^\n]+)\n\n([^\n]+)`)

	// Find all matches in the Markdown content
	matches := pattern.FindAllStringSubmatch(markdownContent, -1)

	// Create an SLO instance
	var slo SLO

	// Iterate over the matches and assign them to the struct
	for _, match := range matches {
		if len(match) == 3 {
			// Clean up the matched strings
			key := strings.TrimSpace(match[1])
			value := strings.TrimSpace(match[2])

			// Assign values to struct fields based on the key
			switch key {
			case "SLO Name":
				slo.Name = value
			case "SLO Description":
				slo.Description = value
			case "Numerator Query":
				slo.NumeratorQuery = value
			case "Denominator Query":
				slo.DenominatorQuery = value
			case "7-day Thresholds":
				slo.SevenDayThresh = value
			case "30-day Thresholds":
				slo.ThirtyDayThresh = value
			}
		}
	}
  
  	// Parse the template file.
	tmpl, err := template.ParseFiles("slo.tmpl")
	if err != nil {
		fmt.Println("error parsing template file:", err)
	}

	// Execute the template and write the output to a file.
	file, err = os.Create("slo.tf")
	if err != nil {
		fmt.Println("error creating Terraform file:", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, slo); err != nil {
		fmt.Println("error executing template:", err)
	}
}

