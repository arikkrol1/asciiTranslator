
package translation

import(
    // "fmt"
    "./io"
)

type translationEngine struct{
    numberProvider io.IInputProvider 
    outputHandler io.IOutputHandler
    translationDic map[string]byte
}

//NewTranslationEngine ctor 
func NewTranslationEngine(config map[string]interface{}) *translationEngine{
    te := &translationEngine{}
    te.numberProvider = io.NewAsciiNumberProvider(config)
    te.outputHandler = io.NewInvoiceNumberOutputHandler(config)
    te.initTranslationDic(config)
    return te;
}

func (te *translationEngine) Translate(){
    defer te.closeIO()
    
    for {
        numRepresentation := te.numberProvider.GetNext()
        if numRepresentation == nil {
            return
        }
        
        num := te.translateNumber(numRepresentation)
        te.outputHandler.HandleOutput(num)
    }
}

func (te *translationEngine) closeIO(){
    te.numberProvider.Close()
    te.outputHandler.Close()
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

func (te *translationEngine) initTranslationDic(config map[string]interface{}){
    te.translationDic = make(map[string]byte)
    
    for key,val := range config["digitTranslation"].(map[string]interface{}){
        te.translationDic[key] = val.(string)[0]
    }
}