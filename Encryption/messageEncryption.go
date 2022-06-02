package main

import (
	"fmt"
)

func main() {
	// message, chiffer and array for encrypted message
	message := []rune(", my message in lower case")
	chiffer := []rune("abcdefghijklmnopqrstuvwxyzæøå, ")
	var shiftedMessage []int32
	caesarShift := 2

	// iterates over runes in message
	for i := 0; i < len(message); i++ {
		messageChar := message[i]

		// looks up in alphabet and caesarShift encrypts
		for n := 0; n < len(chiffer); n++ {
			if messageChar == chiffer[n] {
				encryptedChar := chiffer[(n+caesarShift)%len(chiffer)]
				shiftedMessage = append(shiftedMessage, encryptedChar)
				break
			} else if n == len(chiffer)-1 && message[i] != chiffer[n] {
				fmt.Printf("\nRune: %v, String: '%v'\n", messageChar, string(messageChar))
				panic("Character is not in language, unable to encrypt")
			}
		}

	}

	fmt.Println(message)
	fmt.Println(shiftedMessage)
	fmt.Println(chiffer)
	fmt.Println(caesarShift)

}
