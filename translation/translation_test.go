package translation

import (
    "fmt"
    "testing" 
    "io/ioutil"
    "encoding/json"
)

var config map[string]interface{}

func TestMain(m *testing.M) {
	config = readConfig()
}


func Test_TranslationEngine_translateNumber_readingValidAsciiNumberFromFile_numberParsedCorrectly(t *testing.T){
    trasEngine := NewTranslationEngine(config)
    numRepresentation := trasEngine.numberProvider.GetNext()
    num := trasEngine.translateNumber(numRepresentation)
    
    if num != "600143155"{
        fmt.Println("expected 600143155 but was " + num)
        t.FailNow()    
    }
}


func Test_TranslationEngine_translateNumber_readingInvalidAsciiNumberFromFile_numberParsedWithIllegalSuffix(t *testing.T){
    trasEngine := NewTranslationEngine(config)
    numRepresentation := trasEngine.numberProvider.GetNext()
    num := trasEngine.translateNumber(numRepresentation)
    
    if num != "8192?6248 ILLEGAL"{
        fmt.Println("expected 8192?6248 ILLEGAL but was " + num)
        t.FailNow()    
    }
}

func readConfig() map[string]interface{}{
    configObj := make(map[string]interface{})
    
    fileBytes, err := ioutil.ReadFile("../app.conf")
    if err != nil{
        panic(err)
    }
    
    err = json.Unmarshal(fileBytes, &configObj)
    if err != nil{
        panic(err)
    }
    
    return configObj
}