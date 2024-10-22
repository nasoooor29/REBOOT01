package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintOneLine(font Font, printIt string) error {
	// Get the letters
	var letters []Letter
	for _, letter := range printIt {
		l, err := GetLetter(letter, font)
		if err != nil {
			return err
		}
		letters = append(letters, l)
	}
	// Check the width and height of the letters
	// adjust the width if necessary
	err := CheckLettersWidthAndHeight(letters)
	if err != nil {
		return err
	}
	// Print the letters
	PrintLetters(letters)
	return nil
}

func PrintMultipleLines(font Font, printIt string) {
	// split the input by new line
	splitted := strings.Split(printIt, "\\n")
	isEmpty := true
	count := 0

	// check if the input contains non-ASCII characters
	allASCII := CheckIfAllIsPrintableASCII([]rune(printIt))
	if !allASCII {
		fmt.Println("error: the input contains non-ASCII characters. Please provide only ASCII characters.")
		return
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
	if isEmpty {
		count--
		for i := 0; i < count; i++ {
			fmt.Println("")
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
		err := PrintOneLine(font, line)
		if err != nil {
			return
		}
	}
}

func main() {
	// Get the arguments
	args := os.Args[1:]
	// if there are no arguments, return and print nothing
	if len(args) == 0 {
		return
	}
	// if there are more than one argument, print the usage and return
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Example: go run main.go \"Hello, World!\" shadow")
		return
	}
	// Get the text to print
	toPrint := args[0]
	if len(args) == 1 {
		font, err := GetFont("standard")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		PrintMultipleLines(font, toPrint)
		return
	}
	// Get the standard font
	fontName := args[1]
	font, err := GetFont(fontName)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	PrintMultipleLines(font, toPrint)
}
