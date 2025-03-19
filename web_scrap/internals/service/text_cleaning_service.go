package service

import (
	"strings"
	"unicode"
	"regexp"
	"net/url"
)

func cleanWord(word string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r 
		}
		return -1 
	}, word)
}

func cleanWords(words []string) []string {
	var cleaned []string
	for _, word := range words {
		cleanedWord := cleanWord(word)
		cleanedWord = strings.ToLower(cleanedWord)
		if cleanedWord != "" {
			cleaned = append(cleaned, cleanedWord)
		}
	}
	return cleaned
}

func removeDecimalFromLastNumber(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return ""
	}
	
	re := regexp.MustCompile(`(\d+)\.(\d+)`)
	modifiedPath := re.ReplaceAllString(parsedURL.Path, "$1") 
	parsedURL.Path = modifiedPath
	return parsedURL.String()
}