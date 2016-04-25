package main

import(
    "os"
    "fmt"
    "github.com/arikkrol/asciiTranslator/translation"
)

func main()  {
    defer failCheck()
    
    args := os.Args[1:]
    if len(args) != 2{
        fmt.Println("usage - required args: <input file path> <output file path>")
        return
    }
    
    te := translation.NewTranslationEngine(args[0], args[1])
    te.Translate()
}

func failCheck(){
    if r := recover(); r != nil {
        fmt.Println("Panic caught: ", r)
    }
}
