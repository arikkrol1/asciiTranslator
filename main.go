package main

import(
    "github.com/arikkrol/asciiTranslator/translation"
)

func main()  {
    te := translation.NewTranslationEngine("./translation/io/testdata/good.txt", "./translation/io/testoutput/good.txt")
    te.Translate()
}

