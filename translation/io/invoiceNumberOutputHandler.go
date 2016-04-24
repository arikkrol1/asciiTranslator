package io

import(
    "os"
    "fmt"
)

type invoiceNumberOutputHandler struct{
    outFile *os.File
}

//NewInvoiceNumberOutputHandler ctor
func NewInvoiceNumberOutputHandler(file string) *invoiceNumberOutputHandler{
    handler := &invoiceNumberOutputHandler{}
    handler.Open(file)
    return handler
}

func (handler invoiceNumberOutputHandler) Open(file string){
    var err error 
    handler.outFile, err = os.Open(file)
    if err != nil {
        fmt.Println(err)
        //TODO: use panic
        
        return
    }
}



func (handler invoiceNumberOutputHandler) HandleOutput(output string){
    //TODO: complete
} 