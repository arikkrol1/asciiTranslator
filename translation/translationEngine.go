
package translation

import(
    // "fmt"
    "github.com/arikkrol/asciiTranslator/translation/io"
)


type translationEngine struct{
    numberProvider io.IInputProvider 
    outputHandler io.IOutputHandler
    translationDic map[string]string
}

//NewTranslationEngine ctor 
func NewTranslationEngine(inFile string, outFile string) *translationEngine{
    ta := &translationEngine{}
    
    ta.numberProvider = io.NewAsciiNumberProvider(inFile)
    //TODO: complete outputHandler
    
    //TODO: complete
    return nil;
}