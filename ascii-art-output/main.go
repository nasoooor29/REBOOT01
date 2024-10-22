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

func WriteOneLine(font Font, printIt string) (error, string) {
	// Get the letters
	var letters []Letter
	for _, letter := range printIt {
		l, err := GetLetter(letter, font)
		if err != nil {
			return err, ""
		}
		letters = append(letters, l)
	}
	// Check the width and height of the letters
	// adjust the width if necessary
	err := CheckLettersWidthAndHeight(letters)
	if err != nil {
		return err, ""
	}
	// Print the letters
	return nil, WriteLetters(letters)
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
func WriteMultipleLines(font Font, printIt string) string {
	finalStr := ""
	// split the input by new line
	splitted := strings.Split(printIt, "\\n")
	isEmpty := true
	count := 0

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
			finalStr += "\n"
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
			finalStr += "\n"
			continue
		}

		// print the line
		err, str := WriteOneLine(font, line)
		if err != nil {
			return ""
		}
		finalStr += str
	}
	return finalStr
}

func ValidExtension(fileName string) bool {
	//Minimum filename length is 5. Example: a.txt
	if len(fileName) < 5 {
		return false
	}
	if fileName[len(fileName)-3:] == "txt" {
		return true
	}
	return false
}

func main() {
	// Get the arguments
	args := os.Args[1:]
	// if there are no arguments, return and print nothing
	if len(args) == 0 {
		return
	}
	// if there are more than one argument, print the usage and return
	if len(args) != 3 {
		fmt.Println("Usage: go run main.go <text to print>")
		fmt.Println("Example: go run main.go --output=filename.txt \"Hello, World!\" fontname")
		return
	}

	// Get the text to print
	toPrint := args[1]
	// Get the standard font
	fontName := args[2]
	font, err := GetFont(fontName)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	PrintMultipleLines(font, toPrint)
	fileName := args[0][9:]
	if ValidExtension(fileName) {
		fmt.Println(fileName)
		bytes := []byte(WriteMultipleLines(font, toPrint))
		err = os.WriteFile(fileName, bytes, 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Output filename invalid")
		return
	}

}
