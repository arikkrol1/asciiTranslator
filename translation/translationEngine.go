
package translation

import(
    "fmt"
    "github.com/arikkrol/asciiTranslator/translation/io"
)

type translationEngine struct{
    numberProvider io.IInputProvider 
    outputHandler io.IOutputHandler
    translationDic map[string]byte
}

//NewTranslationEngine ctor 
func NewTranslationEngine(inFile string, outFile string) *translationEngine{
    te := &translationEngine{}
    te.numberProvider = io.NewAsciiNumberProvider(inFile)
    te.outputHandler = io.NewInvoiceNumberOutputHandler(outFile)
    te.initTranslationDic()
    return te;
}

func (te *translationEngine) Translate(){
    
    for {
        numRepresentation := te.numberProvider.GetNext()
        
        if numRepresentation == nil {
            return
        }
        
        num := te.translateNumber(numRepresentation)
        
        //TODO: remove
        fmt.Println(num)
        
        te.outputHandler.HandleOutput(num)
    }
}

func (te *translationEngine) translateNumber(numRepresentation []string) string{
    suffix := ""
    numArr := make([]byte, len(numRepresentation))
    for i,val := range numRepresentation{
        char, exists := te.translationDic[val]
        if exists{
            numArr[i] = char    
        } else {
            numArr[i] = '?'
            suffix = " ILLEGAL"
        }
    }
    
    res := string(numArr) + suffix
    return res
}

func (te *translationEngine) initTranslationDic(){
    te.translationDic = make(map[string]byte)
    
    te.translationDic[" _ | ||_|"] = '0';
    te.translationDic["     |  |"] = '1';
    te.translationDic[" _  _||_ "] = '2';
    te.translationDic[" _  _| _|"] = '3';
    te.translationDic["   |_|  |"] = '4';
    te.translationDic[" _ |_  _|"] = '5';
    te.translationDic[" _ |_ |_|"] = '6';
    te.translationDic[" _   |  |"] = '7';
    te.translationDic[" _ |_||_|"] = '8';
    te.translationDic[" _ |_| _|"] = '9';
}