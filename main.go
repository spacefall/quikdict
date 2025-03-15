package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"os"
)

const apiUrl = "https://api.dictionaryapi.dev/api/v2/entries/en/"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s No word was provided\n", color.HiRedString("Error:"))
		fmt.Printf("Usage: %s [word]", os.Args[0])
		os.Exit(1)
	}

	word := os.Args[1]

	r, err := http.Get(apiUrl + word)
	if err != nil {
		fmt.Printf("%s %s\n", color.HiRedString("Error:"), err)
		os.Exit(1)
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		fmt.Printf("%s %s\n", color.HiRedString("Error:"), "Word not found")
		os.Exit(1)
	}

	// Parse the response body
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s %s\n", color.HiRedString("Error:"), err)
		os.Exit(1)
	}

	value := gjson.ParseBytes(bytes)

	print(value.String())
}
