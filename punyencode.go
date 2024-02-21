package main

import (
	"fmt"
	"golang.org/x/net/idna"
	"os"
)

func main() {
	// Check if the non-ASCII string is provided as command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: punycode_encoder <non-ASCII>")
		return
	}

	// Retrieve the non-ASCII string from command line argument
	nonAscii := os.Args[1]

	// Encode the non-ASCII string into Punycode
	punycode, err := idna.ToASCII(nonAscii)
	if err != nil {
		fmt.Println("Error encoding non-ASCII to Punycode:", err)
		return
	}

	// Print the Punycode representation
	fmt.Println("Punycode:", punycode)
}
