package io

import(
    "os"
    //"fmt"
    "bufio"
    "runtime"
)

type invoiceNumberOutputHandler struct{
    outFile *os.File
    writer *bufio.Writer
    lineFeed string
}

//NewInvoiceNumberOutputHandler ctor
func NewInvoiceNumberOutputHandler(config map[string]interface{}) *invoiceNumberOutputHandler{
    handler := &invoiceNumberOutputHandler{}
    handler.Open(config["outputFile"].(string))
    
    if runtime.GOOS == "windows" {
        handler.lineFeed = "\r\n"
    } else {
        handler.lineFeed = "\n"
    }
    
    return handler
}

func (handler *invoiceNumberOutputHandler) Close(){
    handler.outFile.Close()
}

func (handler *invoiceNumberOutputHandler) Open(file string){
    var err error 
    handler.outFile, err = os.Create(file)
    if err != nil {
        panic("cant open file " + file)
    }
    
    handler.writer = bufio.NewWriter(handler.outFile)
}


func (handler *invoiceNumberOutputHandler) HandleOutput(output string){
    handler.writer.WriteString(output)
    handler.writer.WriteString(handler.lineFeed)
    handler.writer.Flush()
} 