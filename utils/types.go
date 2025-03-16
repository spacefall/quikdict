package utils

type Definition struct {
	Definition string
	Example    string
}

type Meaning struct {
	PartOfSpeech string
	Definitions  []Definition
	Synonyms     []string
	Antonyms     []string
}

type WordInfo struct {
	Word      string
	Phonetics []string
	Meanings  []Meaning
	Sources   []string
}
