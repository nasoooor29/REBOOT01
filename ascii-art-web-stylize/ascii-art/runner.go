package ascii

import (
	"fmt"
	"strings"
)

func PrintOneLine(font Font, printIt string) ([]string, error) {
	// Get the letters
	var letters []Letter
	for _, letter := range printIt {
		l, err := GetLetter(letter, font)
		if err != nil {
			return nil, err
		}
		letters = append(letters, l)
	}
	// Check the width and height of the letters
	// adjust the width if necessary
	err := CheckLettersWidthAndHeight(letters)
	if err != nil {
		return nil, err
	}
	// Print the letters
	PrintLetters(letters)
	var output []string
	for i := 0; i < LETTER_HEIGHT; i++ {
		var line string
		for _, letter := range letters {
			line += letter[i]
		}
		output = append(output, line)
	}
	return output, nil
}

func PrintMultipleLines(font Font, printIt string) ([]string, error) {
	// split the input by new line
	printIt = strings.ReplaceAll(printIt, "\r", "")
	splitted := strings.Split(printIt, "\n")
	isEmpty := true
	count := 0

	// check if the input contains non-ASCII characters
	allASCII := CheckIfAllIsPrintableASCII([]rune(printIt))
	if !allASCII {
		return nil, fmt.Errorf("error: the input contains non-ASCII characters. Please provide only ASCII characters")
	}
	// check if the input is empty
	for _, line := range splitted {
		// if the word is not empty, set isEmpty to false
		if len(line) > 0 {
			isEmpty = false
		} else {
			count++
		}
	}

	// if the input is empty, print new lines
	var output []string
	if isEmpty {
		for i := 0; i < count-1; i++ {
			output = append(output, "")
		}
	}

	// print the input
	for _, line := range splitted {
		// if the is empty flag is true, continue
		if isEmpty {
			continue
		}

		// if the line is empty, print a new line
		if line == "" {
			fmt.Println("")
			continue
		}

		// print the line
		lineOutput, err := PrintOneLine(font, line)
		if err != nil {
			return nil, err
		}
		output = append(output, lineOutput...)
	}
	return output, nil
}

func Output(toPrint, fontName string) (string, error) {


	// load the specified font
	font, err := GetFont(fontName)
	if err != nil {
		return "", fmt.Errorf("error in getting font")
	}

	// generate the output
	output, err := PrintMultipleLines(font, toPrint)
	if err != nil {
		return "", fmt.Errorf("error in printing lines")
	}

	outputString := strings.Join(output, "\n")
	return outputString, nil
}
