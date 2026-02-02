package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Generate(text, banner string) (string, error) {
	// Validate banner
	validBanners := map[string]bool{
		"shadow":     true,
		"standard":   true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {
		return "", fmt.Errorf("invalid banner: %s", banner)
	}

	// Load banner file
	fileName := fmt.Sprintf("banner/%s.txt", banner)
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("banner file not found: %s", fileName)
	}
	defer file.Close()

	// Read banner characters
	scanner := bufio.NewScanner(file)
	bannerChars := make([]string, 0)
	for scanner.Scan() {
		bannerChars = append(bannerChars, scanner.Text())
	}

	// Process text
	lines := strings.Split(text, "\n")
	result := ""

	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}

		// Generate each of the 8 lines
		for i := 0; i < 8; i++ {
			for _, char := range line {
				// Calculate position in banner file
				// ASCII characters start from 32 (space)
				if char < 32 || char > 126 {
					return "", fmt.Errorf("invalid character: %c", char)
				}
				
				pos := int(char-32)*9 + i
				if pos < len(bannerChars) {
					result += bannerChars[pos]
				}
			}
			result += "\n"
		}
	}

	return result, nil
}