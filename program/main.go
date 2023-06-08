package main

import (
	"asciiart"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) < 3 {
		fmt.Println("File name missing")
		return
	}
	file66 := os.Args[1]
	file, err := ioutil.ReadFile(file66)
	if err != nil {
		fmt.Print(err)
		return
	}
	Args := os.Args[2]
	fmt.Println(Args)
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 2)
		file1 := asciiart.ChooseString(file, arrayofrune)
		fmt.Println(string(file1))

	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 3)
		file2 := asciiart.ChooseString(file, arrayofrune)
		fmt.Println(string(file2))

	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 4)
		file3 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file3))

	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 5)
		file4 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file4))
	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 6)
		file5 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file5))
	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 7)
		file6 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file6))

	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 8)
		file7 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file7))
	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 9)
		file8 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file8))

	}
	for _, c := range Args {
		arrayofrune := (((int(c) - 32) * 9) + 10)
		file9 := asciiart.ChooseString(file, arrayofrune)

		fmt.Println(string(file9))
	}
}
