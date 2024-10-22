package main

import (
	"fmt"
	"slices"
)

// CheckIfAllIsPrintableASCII checks if all the characters in the input are printable ASCII characters
func CheckIfAllIsPrintableASCII(chars []rune) bool {
	allASCII := true
	for _, char := range chars {
		if (char >= FIRST_ASCII_CHAR && char <= LAST_ASCII_CHAR) || char == '\n' {
			continue
		}
		allASCII = false
	}
	return allASCII
}

// CheckLettersHight checks if all the letters have the same height
func CheckLettersHight(letters []Letter) error {
	for _, letter := range letters {
		if len(letter) != LETTER_HEIGHT {
			return fmt.Errorf("there is a letter having more or less than 8 char")
		}
	}
	return nil
}

// AdjustWidthSpaces adjusts the width of the letters
// by adding spaces to the right of the letters
// to make all the letters have the same width
func AdjustWidthSpaces(letters []Letter) error {
	for _, letter := range letters {
		m := FindThe(slices.Max, letter)
		for i, VerticalLetter := range letter {
			if len(VerticalLetter) != m {
				diff := m - len(VerticalLetter)
				spaces := GenerateSpaces(diff)
				letter[i] = VerticalLetter + spaces
			}
		}
	}
	return nil
}

// AdjustSpacesBetweenLetters adjusts the spaces between the letters
// by adding spaces to the right of the letters
func AdjustSpacesBetweenLetters(letters []Letter) {
	for letterIdx, letter := range letters {
		for i, VerticalLetter := range letter {
			if letterIdx == len(letters)-1 {
				spaces := GenerateSpaces(1)
				letter[i] = VerticalLetter + spaces
				continue
			}

			if len(VerticalLetter) == 0 {
				spaces := GenerateSpaces(SPACE_WIDTH)
				letter[i] = VerticalLetter + spaces
				continue
			}
			spaces := GenerateSpaces(SPACE_BETWEEN_LETTERS)
			letter[i] = VerticalLetter + spaces
		}
	}
}

// just a helper function to find the max number in a slice
func FindThe(fn func([]int) int, letter Letter) int {
	numbers := []int{}
	for _, verticalLetter := range letter {
		numbers = append(numbers, len(verticalLetter))
	}
	return fn(numbers)
}

// GenerateSpaces generates spaces based on the param
func GenerateSpaces(num int) string {
	str := ""
	for range num {
		str += " "
	}
	return str
}

// CheckLettersWidthAndHeight checks the width and height of the letters
// the function adjusts the width and spaces between the letters if necessary
func CheckLettersWidthAndHeight(letters []Letter) error {
	err := CheckLettersHight(letters)
	if err != nil {
		return err
	}
	AdjustWidthSpaces(letters)
	AdjustSpacesBetweenLetters(letters)
	return nil
}
