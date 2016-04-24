
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
    te.initTranslationDic()
    return te;
}

func (te translationEngine) Translate(){
    //TODO: complete
}

func (te translationEngine) initTranslationDic(){
    te.translationDic = make(map[string]string)
    
    te.translationDic[" _ | ||_|"] = "0";
    te.translationDic["     |  |"] = "1";
    te.translationDic[" _  _||_ "] = "2";
    te.translationDic[" _  _| _|"] = "3";
    te.translationDic["   |_|  |"] = "4";
    te.translationDic[" _ |_  _|"] = "5";
    te.translationDic[" _ |_ |_|"] = "6";
    te.translationDic[" _   |  |"] = "7";
    te.translationDic[" _ |_||_|"] = "8";
    te.translationDic[" _ |_| _|"] = "9";
}