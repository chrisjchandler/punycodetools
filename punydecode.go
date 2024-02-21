package main

import (
    "fmt"
    "golang.org/x/net/idna"
    "os"
)

func main() {
    // Check if the punycode string is provided as command line argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: punycode_decoder <punycode>")
        return
    }

    // Retrieve the punycode string from command line argument
    punycode := os.Args[1]

    // Decode the punycode string
    nonAscii, err := idna.ToUnicode(punycode)
    if err != nil {
        fmt.Println("Error decoding punycode:", err)
        return
    }

    // Print the decoded non-ASCII string
    fmt.Println("Non-ASCII string:", nonAscii)
}
