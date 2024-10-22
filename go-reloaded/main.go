package main

import (
	"fmt"
	"os"
)

func ReadArgs() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("the app should have 3 args only")
	}
	return nil
}

func ReadFileContents(fileName string) ([]byte, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func WriteFileContents(fileName string, data string) error {
	// create file
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error creating file:", err)
		return err
	}
	// remember to close the file
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := ReadArgs()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Println("usage: ./main <input file name> <output file name>")
		return
	}
	in := os.Args[1]
	out := os.Args[2]
	data, err := ReadFileContents(in)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	result, err := ReloadedWithNewLine(string(data))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	err = WriteFileContents(out, result)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}
