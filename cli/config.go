package cli

import (
	"fmt"
	"os"
)

var HelpStr = fmt.Sprintf(`Usage: %s [word] [options]
Options:
  -th, --thesaurus    Print synonyms and antonyms
  -h, --help          Print this help message`, os.Args[0])

type Params struct {
	Word      string
	Thesaurus bool
}
