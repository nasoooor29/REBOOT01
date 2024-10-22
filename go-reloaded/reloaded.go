package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type EditLastNFunc func(arrPtr *[]string, from, to int) error
type EditorFunc func(string) string

func CompareBrackets(str string) bool {
	rCounter := 0
	lCounter := 0
	for _, v := range str {
		if v == '(' {
			lCounter++
		}
		if v == ')' {
			rCounter++
		}
	}
	return rCounter == lCounter
}

func RemoveNonAscii(str string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	return re.ReplaceAllString(str, "")
}

func EditLastNWords(s *[]string, fn EditorFunc, from, to int) error {
	num := to - from
	if from < 0 || to > len(*s) {
		return fmt.Errorf("invalid range")
	}
	c := 0
	for i := to - 1; i >= from; i-- {
		w := &((*s)[i])
		arr := strings.Split(*w, ",")
		for j := len(arr) - 1; j >= 0; j-- {
			if c >= num {
				break
			}
			arr[j] = fn(arr[j])
			c++
		}
		*w = strings.Join(arr, ",")
	}
	return nil
}

func curryFunc(fn func(string) string) EditLastNFunc {
	return func(arrPtr *[]string, from, to int) error {
		return EditLastNWords(arrPtr, fn, from, to)
	}
}

func HexToInt(hex string) string {
	number, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return "HEX ERROR"
	}
	return fmt.Sprint(number)
}

func BinToInt(bin string) string {
	number, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return "BIN ERROR"
	}
	return fmt.Sprint(number)
}

var Mapping = map[string]EditLastNFunc{
	"up":  curryFunc(strings.ToUpper),
	"low": curryFunc(strings.ToLower),
	"cap": curryFunc(strings.Title),
	"hex": curryFunc(HexToInt),
	"bin": curryFunc(BinToInt),
}

type modifierWord struct {
	word string
	num  int
}

func NewModifierWord(s string) (*modifierWord, error) {
	var word string
	var num int
	var err error
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "(", "")
	if strings.Contains(s, ",") {
		split := strings.Split(s, ",")
		word = split[0]
		num, err = strconv.Atoi(split[1])
		if err != nil {
			return nil, fmt.Errorf("invalid modifier")
		}
	} else {
		word = s
		num = 1
	}
	return &modifierWord{word, num}, nil
}

func SeparateTheBracket(s string) string {
	// word( -> word(
	s = regexp.MustCompile(`(\w+)\(`).ReplaceAllString(s, "${1} (")
	// )word -> ) word
	s = regexp.MustCompile(`\)(\w+)`).ReplaceAllString(s, ") ${1}")
	return s
}

func RemoveDoubleSpaces(s string) string {
	r := regexp.MustCompile(`\s+`)
	return r.ReplaceAllString(s, " ")
}

func EditModifiersSpaces(s string) string {
	s = strings.ReplaceAll(s, "(", "( ")
	s = strings.ReplaceAll(s, ")", " )")
	s = strings.ReplaceAll(s, ",", " , ")
	return s
}

func EditModifiersToCorrectFormat(s string) string {
	re := regexp.MustCompile(`\(\s(\w+)\s,\s(\d+)\s\)`)
	s = re.ReplaceAllString(s, "($1,$2)")
	re = regexp.MustCompile(`\(\s*(\w+)\s*\)`)
	s = re.ReplaceAllString(s, "($1,1)")
	return s
}
func NormalReloaded(in string) (string, error) {
	in = RemoveNonAscii(in)
	b := CompareBrackets(in)
	if !b {
		return "", fmt.Errorf("not every bracket has a pair")
	}
	in = SeparateTheBracket(in)
	// ([spaces]word[spaces],[spaces] number[spaces]) -> ( word , number )
	in = EditModifiersSpaces(in)
	// "  " -> " "
	in = RemoveDoubleSpaces(in)
	// (word) -> (word,1)
	// ([spaces]word[spaces],[spaces] number[spaces]) -> (word,number)
	// ([spaces]word[spaces]) -> (word,1)
	in = EditModifiersToCorrectFormat(in)

	in = strings.TrimSpace(in)
	in = SeparateBrackets(in)
	in = regexp.MustCompile(`\b(\w+),\s*(\d+)\b`).ReplaceAllString(in, "${1},${2}")
	words := strings.Split(in, " ")
	words = removeEmptyStrings(words)

	for range len(in) {
		err := ResolveEverything(&words)
		if err != nil {
			return "", err
		}
	}

	words = removeEmptyStrings(words)
	ResolveQuotes(&words)
	str := strings.Join(words, " ")
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "( ", "(")
	str = strings.ReplaceAll(str, ") ", ")")

	return strings.ReplaceAll(str, "  ", " "), nil
}

func ReloadedWithNewLine(in string) (string, error) {
	arr := strings.Split(in, "\n")
	var res []string
	for _, v := range arr {
		s, err := NormalReloaded(v)
		if err != nil {
			return "", err
		}
		res = append(res, s)
	}
	return strings.Join(res, "\n"), nil
}

func ResolveEverything(wordsPtr *[]string) error {

	for i := 0; i < len(*wordsPtr); i++ {
		word := (*wordsPtr)[i]
		if len(word) == 0 {
			continue
		}
		s, err := ResolveModifierWord(i, word, wordsPtr)
		if err != nil {
			return err
		}
		if s {
			i--
		}
		ResolveA(wordsPtr, i)
		ResolvePunc(i, word, wordsPtr)
	}
	return nil
}

func removeEmptyStrings(s []string) []string {
	var res []string
	for _, v := range s {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func ResolvePunc(i int, word string, s *[]string) {
	puncs := map[string]int{",": 1, ".": 1, "!": 1, "?": 1, ":": 1, ";": 1}
	modifierWordRegex := regexp.MustCompile(`\([^()]*?(,\s*\d+)?\)`)
	PuncsRegex := regexp.MustCompile(`[",.!?;:]`)
	thereIsPunc := PuncsRegex.MatchString(word)
	modifierWord := modifierWordRegex.MatchString(word)
	if !thereIsPunc || modifierWord {
		return
	}

	if len(*s) == 1 {
		return
	}

	if i-1 < 0 {
		return
	}

	if word == "..." || word == "!?" {
		(*s)[i-1] += word
		(*s)[i] = ""
		return
	}
	// if there is a punctuation on the middle of the array merge it with the word before
	if _, foundPunc := puncs[word]; foundPunc {
		(*s)[i-1] += word
		if i == len(*s)-1 {
			(*s)[len(*s)-1] = ""
			return
		}
		(*s)[i] = ""
		return
	}

	// if there is a word with punctuation at the beginning split it
	if _, foundPunc := puncs[string(word[0])]; foundPunc {
		(*s)[i-1] += string(word[0])
		(*s)[i] = word[1:]
	}
}

func ResolveModifierWord(i int, word string, words *[]string) (bool, error) {
	r := regexp.MustCompile(`\([^()]*?(,\s*\d+)?\)`)
	lowerWord := strings.ToLower(word)
	modifier := r.MatchString(lowerWord)
	if !modifier {
		return false, nil
	}
	w, err := NewModifierWord(lowerWord)
	if err != nil {
		return false, fmt.Errorf("invalid modifier word: %v", word)
	}
	fn, ok := Mapping[w.word]
	if !ok {
		return false, fmt.Errorf("invalid modifier word: (%v)", w.word)
	}
	err = fn(words, i-w.num, i)
	if err != nil {
		return false, err
	}
	*words = append((*words)[:i], (*words)[i+1:]...)
	return true, nil
}

func ResolveA(s *[]string, currentPos int) {
	vowels := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'h': 1}
	if currentPos <= len(*s)-2 {
		currentWord := (*s)[currentPos]
		if strings.ToLower(currentWord) != "a" {
			return
		}
		nextWord := strings.ToLower((*s)[currentPos+1])
		nextWordFirstLetter := rune(nextWord[0])
		if currentWord == "a" && vowels[nextWordFirstLetter] == 1 {
			(*s)[currentPos] = "an"
		}
		if currentWord == "A" && vowels[nextWordFirstLetter] == 1 {
			(*s)[currentPos] = "An"
		}
	}
}

func ResolveQuotes(words *[]string) {
	// for apostrophe
	count := 0
	for i, word := range *words {
		if word == "'" && count == 0 {
			count += 1
			(*words)[i+1] = word + (*words)[i+1]
			*words = append((*words)[:i], (*words)[i+1:]...)
		}
	}
	//  for second apostrophe
	for i, word := range *words {
		if word == "'" {
			(*words)[i-1] = (*words)[i-1] + word
			*words = append((*words)[:i], (*words)[i+1:]...)
		}
	}
}

func SeparateBrackets(txt string) string {
	re := regexp.MustCompile(`(\w+)\s*(\([^)]*\))(\w+)`)
	result := re.ReplaceAllString(txt, "$1 $2 $3")
	return result
}

func CompareArrays(a1, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true

}
