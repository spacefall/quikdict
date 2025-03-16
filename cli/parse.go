package cli

import (
	"fmt"
	"os"
)

func Parse(args []string) Params {
	var params Params
	for _, arg := range args {
		switch arg {
		case "-th", "--thesaurus":
			params.Thesaurus = true

		case "-h", "--help":
			fmt.Println(HelpStr)
			os.Exit(0)

		default:
			params.Word = arg
		}
	}
	return params
}
