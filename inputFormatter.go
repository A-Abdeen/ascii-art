package asciiart

import "strings"

func InputFormatter(rawInput string) string {
	// TODO add to README instructions
	formattedText := strings.ReplaceAll(rawInput, "\\n", "\n")
	formattedText = strings.ReplaceAll(formattedText, "\"", string('"'))
	formattedText = strings.ReplaceAll(formattedText, "\\'", "'")
	return formattedText
}
