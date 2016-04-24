package io

import(
    "os"
    "fmt"
    "bufio"
)

type invoiceNumberOutputHandler struct{
    outFile *os.File
    writer *bufio.Writer
}

//NewInvoiceNumberOutputHandler ctor
func NewInvoiceNumberOutputHandler(file string) *invoiceNumberOutputHandler{
    handler := &invoiceNumberOutputHandler{}
    handler.Open(file)
    return handler
}

func (handler *invoiceNumberOutputHandler) Close(){
    handler.outFile.Close()
}

func (handler *invoiceNumberOutputHandler) Open(file string){
    var err error 
    handler.outFile, err = os.Open(file)
    if err != nil {
        fmt.Println(err)
        //TODO: use panic
        
        return
    }
    
    handler.writer = bufio.NewWriter(handler.outFile)
}



func (handler *invoiceNumberOutputHandler) HandleOutput(output string){
    handler.writer.WriteString(output)
    //TODO: maybe need to write linefeed
} 