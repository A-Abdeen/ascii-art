package asciiart

func ChooseString (standard []byte, valueofrune int) []byte {
	byteofchar := []byte{}
	count := 1
	j := 1
	for i := 0;i >= 0;i++ {
	if standard[i] == 10 {
		count++
	}
	if count == valueofrune {
		j = i + 1
		break
	}
	// testing
	}
	for ;j > 0; j++ {
	if standard[j] == 10 {
		break
	}
			byteofchar = append(byteofchar, standard[j])
	}
	return byteofchar
}