package main

import(
    "os"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/arikkrol/asciiTranslator/translation"
)

func main()  {
    defer failCheck()
    
    args := os.Args[1:]
    if len(args) != 2{
        fmt.Println("usage - required args: <input file path> <output file path>")
        return
    }
    
    config := readConfig()
    fmt.Println(config)
    
    te := translation.NewTranslationEngine(args[0], args[1])
    te.Translate()
}

func failCheck(){
    if r := recover(); r != nil {
        fmt.Println("Panic caught: ", r)
    }
}

func readConfig() map[string]interface{}{
    configObj := make(map[string]interface{})
    
    fileBytes, err := ioutil.ReadFile("./app.conf")
    if err != nil{
        panic(err)
    }
    
    err = json.Unmarshal(fileBytes, &configObj)
    if err != nil{
        panic(err)
    }
    
    return configObj
}