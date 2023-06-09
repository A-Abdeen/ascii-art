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

	// First function: Modifier (change input to allow use of newline & qoutation marks)
	formattedInput := asciiart.InputFormatter(rawInput)

	// Second function: Splitting (split string)
	stringLine := ""
	var testString string
	var multiString []string
	for x, y := range formattedInput {

		if y == 10 {
			multiString = append(multiString, testString)
			testString = ""
		} else {
			testString = testString + string(y)
		}
		if x == len(formattedInput)-1 && testString != "" {
			multiString = append(multiString, testString)
		}
	}

	// Third & Fourth Functions: Parsing & Printing
	for _, args2 := range multiString {
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
