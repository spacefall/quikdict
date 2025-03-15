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
	Phonetics []string
	Meanings  []Meaning
	Sources   []string
}
