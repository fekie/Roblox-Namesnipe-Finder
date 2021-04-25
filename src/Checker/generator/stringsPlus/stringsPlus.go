package stringsPlus

import (
	"bytes"
)

func ReplaceChar(rawString string, newChar byte, location int) string {
	var chars = []byte{}
	for i := 0; i < len(rawString); i++ {
		if i != location {
			chars = append(chars, rawString[i])
		} else {
			chars = append(chars, newChar)
		}
	}
	return string(chars)
}

func InsertChar(rawString string, newChar byte, location int) string {
	var chars = []byte{}
	inserted := false
	for i := 0; i < (len(rawString) + 1); i++ {
		if i != location {
			if !inserted { 
				chars = append(chars, rawString[i])
			} else {
				chars = append(chars, rawString[i - 1])
			}
		} else {
			inserted = true
			chars = append(chars, newChar)
		}
	}
	return string(chars)
}

func CreateFlippedLICombos(rawName string) []string {
	liPositions := []int{}
	lower := bytes.ToLower([]byte(rawName))
	for i := 0; i < len(rawName); i++ {
		char := lower[i]
		if char == 'i' || char == 'l' {
			liPositions = append(liPositions, i)
		} 
	}
	
	fCombos := []string{}
	for i := 0; i < len(liPositions); i++ {
		charPos := liPositions[i]
		char := rawName[charPos]
		var newChar byte
		switch char {
		case 'i', 'I': {
			newChar = 'l'
		}
		case 'l', 'L': {
			newChar = 'i'
		}

		}
		newName := ReplaceChar(rawName, newChar, charPos)
		fCombos = append(fCombos, newName)
	}
	return fCombos
}