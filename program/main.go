package main

import (
	"asciiart"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Check if input is correct
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) < 3 {
		fmt.Println("File name missing")
		return
	}
	// Read char file & string argument
	file := os.Args[1]
	sourceFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	rawInput := os.Args[2]

	// Sub function: Formatting (change input to allow use of newline & qoutation marks)
	// Main function: Splitting (split string based on newline position)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Third & Fourth Functions: Parsing & Printing
	stringLine := ""
	for _, args2 := range splitInput {
		if args2 != "" {
			for i := 2; i < 10; i++ {
				for _, data := range args2 {
					sourceline := (((int(data) - 32) * 9) + i)
					charRow := asciiart.ChooseString(sourceFile, sourceline)
					stringLine = stringLine + string(charRow)
				}
				fmt.Println(stringLine)
				stringLine = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
}
