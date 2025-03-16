package cli

func Parse(args []string) Params {
	var params Params
	for _, arg := range args {
		if arg == "-t" {
			params.Thesaurus = true
		} else {
			params.Word = arg
		}
	}
	return params
}
