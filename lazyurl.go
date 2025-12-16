package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func extractID(url string) (string, string, string, error) {
	// Define patterns for each site and ID type
	patterns := []struct {
		site  string
		param string
		regex string
	}{
		{"Booking.com", "dest_id", `(?:&|\?)dest_id=([^&]+)`},
		{"Hotels.com", "regionId", `(?:&|\?)regionId=([^&]+)`},
		{"Tripadvisor", "geo", `/(g\d+)-`}, // captures the gxxxxxx pattern
	}

	for _, p := range patterns {
		re, err := regexp.Compile(p.regex)
		if err != nil {
			return "", "", "", fmt.Errorf("failed to compile regex for %s: %v", p.param, err)
		}
		matches := re.FindStringSubmatch(url)
		if len(matches) >= 2 {
			return p.site, p.param, matches[1], nil
		}
	}

	return "", "", "", fmt.Errorf("no known ID found (dest_id, regionId, or geo)")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=================================")
	fmt.Println("  Universal Travel URL ID Extractor")
	fmt.Println("=================================")
	fmt.Println("Supports:")
	fmt.Println(" - Booking.com  → dest_id")
	fmt.Println(" - Hotels.com   → regionId")
	fmt.Println(" - Tripadvisor  → geo (gXXXXXX)")
	fmt.Println("Type 'quit' or 'exit' to stop")
	fmt.Println()

	for {
		fmt.Print("Enter URL: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		url := strings.TrimSpace(input)
		if strings.ToLower(url) == "quit" || strings.ToLower(url) == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		if url == "" {
			fmt.Println("Please enter a valid URL")
			fmt.Println()
			continue
		}

		site, param, idValue, err := extractID(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("✓ %s → Found %s: %s\n", site, param, idValue)
		}
		fmt.Println()
	}
}
