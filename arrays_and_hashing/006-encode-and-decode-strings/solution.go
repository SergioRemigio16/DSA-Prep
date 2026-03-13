package main

import (
	"strconv"
	"strings"
)

// Encoded string will look like:
// Int#StringInt#StringInt#String!
// Int - length of the string to read
// # - Separate Int from String
// String - all of the characters that need are gonna be read
// ! - End of message
func encode(strs []string) string {
	var encodedMessage strings.Builder
	for _, message := range strs {
		stringLength := len(message)
		encodedMessage.WriteString(strconv.Itoa(stringLength))
		encodedMessage.WriteString("#")
		encodedMessage.WriteString(message)
	}
	encodedMessage.WriteString("!")
	return encodedMessage.String()
}

func decode(encodedMessages string) []string {
	decodedMessage := []string{}
	// Index position of encodedMessage
	i := 0
	for {
		if encodedMessages[i] != '!' { // If didn't reach the end
			messageLength := ""
			messageLengthInt := 0
			// Read integer runes that make up one number
			for {
				if encodedMessages[i] != '#' {
					messageLength += string(encodedMessages[i])
					i++
				} else {
					messageLengthInt, _ = strconv.Atoi(messageLength)
					break
				}
			}
			// Skip over #
			i++

			var message strings.Builder
			for j := 0; j < messageLengthInt; j++ {
				message.WriteString(string(encodedMessages[i]))
				i++
			}
			decodedMessage = append(decodedMessage, message.String())
		} else {
			return decodedMessage
		}
	}
}

func main() {
	dummy_input := []string{"Hello", "World"}
	println(decode(encode(dummy_input))[0])
	println(decode(encode(dummy_input))[1])
}
