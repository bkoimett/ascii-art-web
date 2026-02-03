package asciiart

import (
	"fmt"
	"os"
	"strings"
	"log"
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

	if text == "\n" {
		return text, nil
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	// text = strings.ReplaceAll(text, "\\r", "")

	words := strings.Split(text, "\\n")

	fileName := fmt.Sprintf("banner/%s.txt", banner)

	banners := LoadBanner(fileName)

	output := PrintAscii(banners, words)

	return output, nil
}

	// Load banner file
func LoadBanner(filename string) map[rune][]string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	if info.Size() < 10 {
		return nil
	}


	lines := strings.Split(string(data), "\n")

	bannerMap := make(map[rune][]string)

	for ch := 32 ; ch <= 126; ch++ {
		start := int(ch -32) * 9
		bannerMap[rune(ch)] = lines[start + 1 : start + 9 ]
	}

	return bannerMap

}


func PrintAscii(banner map[rune][]string, words []string) string {
	result := ""
	
	for _, word := range words {
		if word == ""  || word == " " || word == "\n" {
			result += "\n"
			continue
		}

		word = strings.TrimSpace(word)
	
		for i := 0; i <= 7; i++ {
			for j := 0; j<len(word);j++ {
				char := word[j]
				result += banner[rune(char)][i]
			}
			result += "\n"
		}
	}
	return result
}
