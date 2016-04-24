
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
    te := &translationEngine{}
    te.numberProvider = io.NewAsciiNumberProvider(inFile)
    te.outputHandler = io.NewInvoiceNumberOutputHandler(outFile)
    
    
    //TODO: bootstrap translation dic
    
    return te;
}

func (te translationEngine) Translate(){
    //TODO: complete
}