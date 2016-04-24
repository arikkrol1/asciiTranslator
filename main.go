package main

import(
    "github.com/arikkrol/asciiTranslator/translation"
)

func main()  {
    te := translation.NewTranslationEngine("./translation/io/testdata/bad.txt", "./translation/io/testoutput/bad.txt")
    te.Translate()
}

