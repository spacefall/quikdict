package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"quikdict/cli"
	"quikdict/modes"
	"quikdict/sources"
)

func main() {
	color.NoColor = false
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [word]", os.Args[0])
		handleErrStr("no word was provided")
	}

	args := cli.Parse(os.Args[1:])

	info, err := sources.GetFromDictionaryAPI(args.Word)
	handleErr(err)

	// Print synonyms and antonyms if thesaurus flag is set
	if args.Thesaurus {
		modes.PrintThesaurus(info)
	} else {
		modes.PrintDictionary(info)
	}
}

func handleErrStr(err string) {
	fmt.Printf("%s %s\n", color.HiRedString("Error:"), err)
	os.Exit(1)
}

func handleErr(err error) {
	if err != nil {
		handleErrStr(err.Error())
	}
}
