package main

import(
    "fmt"
    "encoding/json"
    "io/ioutil"
    "./translation"
)

func main()  {
    defer failCheck()
    
    config := readConfig()
    te := translation.NewTranslationEngine(config)
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