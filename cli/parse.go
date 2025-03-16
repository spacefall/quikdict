package cli

import (
	"fmt"
	"os"
)

func Parse(args []string) Params {
	var params Params
	for _, arg := range args {
		switch arg {
		case "-tr", "--translate":
			params.Translate = true
			params.Thesaurus = false

		case "-th", "--thesaurus":
			params.Thesaurus = true
			params.Translate = false

		case "-h", "--help":
			fmt.Println(HelpStr)
			os.Exit(0)

		default:
			params.Word = arg
		}
	}
	return params
}
