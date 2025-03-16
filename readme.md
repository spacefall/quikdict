# quikdict

A simple go program to quickly look up (english) words and get IPA, meanings and sometimes examples, synonyms and antonyms.  
The application can also translate a word to a different language using Google Translate.  

## Usage
You can find a compiled binary in the [releases](https://github.com/spacefall/quikdict/releases).  
You can also compile it yourself by running `go build` after cloning the repo with `git clone https://github.com/spacefall/quikdict.git`.  

After obtaining a binary, you can get the meaning and examples by running:
```bash
./quickdict <word>
```

If you want synonyms and antonyms, you can run:
```bash
./quickdict -th <word>
```

If instead you want to translate a word, run:
```bash
./quickdict -tr <lc> <word>
```
where `<lc>` is the language code of the language you want to translate to, for example: `en` for English, `it` for Italian, `es` for Spanish, `fr` for French, `de` for German.
The source language is obtained automatically by Google Translate.

## Sources
The dictionary/thesaurus data is obtained from [DictionaryAPI](https://dictionaryapi.dev/), which obtains its data from [Wikitionary](https://en.wiktionary.org/wiki/Wiktionary:Main_Page).  
The translation data is obtained from Google Translate.