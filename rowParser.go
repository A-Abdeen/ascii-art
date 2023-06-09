package asciiart

func RowParser(sourceFile []byte, rowLocation int) []byte {
	charRowData := []byte{}
	count := 1
	j := 1
	for i := 0; i >= 0; i++ {
		if sourceFile[i] == 10 {
			count++
		}
		if count == rowLocation {
			j = i + 1
			break
		}
	}
	for ; j > 0; j++ {
		if sourceFile[j] == 10 {
			break
		}
		charRowData = append(charRowData, sourceFile[j])
	}
	return charRowData
}
