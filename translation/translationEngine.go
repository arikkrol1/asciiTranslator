
package translation

import(
    // "fmt"
    "github.com/arikkrol/asciiTranslator/translation/io"
)


type translationEngine struct{
    numberProvider *io.IInputProvider 
    outputHandler *io.IOutputHandler
    translationDic map[string]string
}

//NewTranslationEngine ctor 
func NewTranslationEngine(*io.IInputProvider, *io.IOutputHandler) *translationEngine{
    
    
    //TODO: complete
    return nil;
}