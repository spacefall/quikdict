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
	if len(os.Args) < 2 {
		handleErrStr("no word was provided")
	}

	args, err := cli.Parse(os.Args[1:])
	handleErr(err)

	switch args.Mode {
	case modes.ModeDictionary:
		info, err := sources.GetFromDictionaryAPI(args.Word)
		handleErr(err)
		modes.PrintDictionary(info)
	case modes.ModeThesaurus:
		info, err := sources.GetFromDictionaryAPI(args.Word)
		handleErr(err)
		modes.PrintThesaurus(info)
	case modes.ModeTranslate:
		tran, og, err := sources.Translate(args.Word, args.Translation)
		handleErr(err)
		modes.PrintTranslation(args.Word, tran, og, args.Translation)
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
