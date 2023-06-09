package asciiart

import "fmt"

func RowPrinter(splitInput []string, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	for _, singleLine := range splitInput {
		if singleLine != "" {
			for i := 2; i < 10; i++ {
				for _, charRune := range singleLine {
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					fullRowData = fullRowData + string(charRowData)
				}
				fmt.Println(fullRowData)
				fullRowData = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
}
