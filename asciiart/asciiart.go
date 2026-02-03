package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Generate(text, banner string) (string, error) {
	validBanners := map[string]bool{
		"shadow":     true,
		"standard":   true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {
		return "", fmt.Errorf("invalid banner: %s", banner)
	}

	if text == "" {
		return "", nil
	}

	fileName := fmt.Sprintf("banner/%s.txt", banner)
	banners, err := LoadBanner(fileName)
	if err != nil {
		return "", err
	}

	if banners == nil || len(banners) == 0 {
		return "", fmt.Errorf("failed to load banner: %s", banner)
	}

	output := PrintAscii(banners, text)
	return output, nil
}

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read banner file: %w", err)
	}

	info, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot get file info: %w", err)
	}

	if info.Size() < 10 {
		return nil, fmt.Errorf("banner file is too small: %s", filename)
	}

	lines := strings.Split(string(data), "\n")

	bannerMap := make(map[rune][]string)

	if len(lines) < 855 {
		return nil, fmt.Errorf("banner file is incomplete: %s", filename)
	}

	for ch := 32; ch <= 126; ch++ {
		start := (ch - 32) * 9
		if start+8 >= len(lines) {
			return nil, fmt.Errorf("banner file format error: missing character %q", rune(ch))
		}
		bannerMap[rune(ch)] = lines[start+1 : start+9]
	}

	return bannerMap, nil
}

func PrintAscii(banner map[rune][]string, text string) string {
	lines := strings.Split(text, "\n")
	
	var result strings.Builder
	
	for lineIndex, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		
		for _, char := range line {
			if char < 32 || char > 126 {
				continue
			}
		}
		
		for i := 0; i < 8; i++ {
			for j := 0; j < len(line); j++ {
				char := rune(line[j])
				if char >= 32 && char <= 126 {
					if asciiLines, ok := banner[char]; ok && i < len(asciiLines) {
						result.WriteString(asciiLines[i])
					}
				}
			}
			result.WriteString("\n")
		}
		
		if lineIndex < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	
	return result.String()
}