package translation

import (
    "fmt"
    "testing" 
)


func TestTranslateNumber(t *testing.T){
    trasEngine := NewTranslationEngine("./io/testdata/good.txt", "./io/testoutput/good.txt")
    
    numRepresentation := trasEngine.numberProvider.GetNext()
    
    num := trasEngine.translateNumber(numRepresentation)
    
    if num != "600143155"{
        fmt.Println("expected 600143155 but was " + num)
        t.FailNow()    
    }
}