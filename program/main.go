package main

import (
	"asciiart"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) < 3 {
		fmt.Println("File name missing")
		return
	}
	file := os.Args[1]
	sourceFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	Args := os.Args[2]
	Args = strings.ReplaceAll(Args, "\\n", "\n")
	Args = strings.ReplaceAll(Args, "\"", string('"'))
	Args = strings.ReplaceAll(Args, "\\'", "'")


	// fmt.Println(Args)
	stringLine := ""
	var testString string
	var multiString []string
	for x, y := range Args {

		if y == 10 {
			multiString = append(multiString, testString)
			testString = ""
		} else {
			testString = testString + string(y)
		}
		if x == len(Args)-1 && testString != ""{
			multiString = append(multiString, testString)
		}
	}
	// fmt.Println(len(multiString))
	// fmt.Print(multiString[0])

// count := 0
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
		}else {
			// count++
				fmt.Print("\n")
			}
		}
		// fmt.Println(count)
	}



























