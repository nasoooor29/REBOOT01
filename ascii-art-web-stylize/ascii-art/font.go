package ascii

import (
	"fmt"
	"os"
	"strings"
)

// letter is a slice of strings, each string represents a line of the letter
type Letter []string

// font is a slice of strings, each string represents all lines of a letter
type Font []string

// Constants
const (

	// Font constants
	LETTER_HEIGHT              = 8
	FIRST_ASCII_CHAR           = 32
	LAST_ASCII_CHAR            = 126
	TOTAL_NUMBER_OF_FONT_CHARS = 95
	SPACE_BETWEEN_LETTERS      = 1
	SPACE_WIDTH                = 3

	// Font paths
	SHADOW_FONT_PATH     = "./fonts/shadow.txt"
	STANDARD_FONT_PATH   = "./fonts/standard.txt"
	THINKERTOY_FONT_PATH = "./fonts/thinkertoy.txt"
)

// GetFont returns a font based on the font name
func GetFont(fontName string) (Font, error) {
	fontPath := ""
	// Load the font
	switch strings.ToLower(fontName) {
	case "shadow":
		fontPath = SHADOW_FONT_PATH
	case "standard":
		fontPath = STANDARD_FONT_PATH
	case "thinkertoy":
		fontPath = THINKERTOY_FONT_PATH
	default:
		return nil, fmt.Errorf("font not found")
	}

	// Read the font file
	b, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, fmt.Errorf("could not load the font")
	}
	// Split the font file into lines
	lines := strings.Split(string(b), "\n")
	return lines, nil
}

// GetLetter returns a letter based on the letter and font
// the function returns an error if the letter is not found in the font
func GetLetter(letter rune, font Font) (Letter, error) {
	result := []string{}
	letter = letter - FIRST_ASCII_CHAR
	start := letter * LETTER_HEIGHT
	end := start + LETTER_HEIGHT
	if end > TOTAL_NUMBER_OF_FONT_CHARS*LETTER_HEIGHT {
		return nil, fmt.Errorf("the font does not contain the %v letter", string(letter+FIRST_ASCII_CHAR))
	}
	if start < 0 {
		return nil, fmt.Errorf("the font does not contain the %v letter", string(letter+FIRST_ASCII_CHAR))
	}
	if end < 0 {
		return nil, fmt.Errorf("the font does not contain the %v letter", string(letter+FIRST_ASCII_CHAR))
	}

	for i := start; i < end; i++ {
		result = append(result, font[i])
	}
	return result, nil
}

func PrintLetters(letters []Letter) {
	for i := 0; i < LETTER_HEIGHT; i++ {
		for _, letter := range letters {
			fmt.Printf("%v", letter[i])
		}
		fmt.Println("")
	}
}
