package asciiart

// import "fmt"

func Converttorune(s string) rune{
	d := '0'
	for _, c := range s {
	d = c
	break	
		// c = c + 400
		// // fmt.Println(string(c))
		// fmt.Println(string(122))
	}
	 return d
}
