package main

import (
	"asciiart"
	"fmt"
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
	sourceFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	rawInput := os.Args[2]

	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	asciiart.RowPrinter(splitInput, sourceFile, asciiart.RowParser)

}
