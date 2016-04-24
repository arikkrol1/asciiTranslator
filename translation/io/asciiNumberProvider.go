package io

import (
    // "os"
    // "fmt"
    // "bufio"
)

type asciiNumberProvider struct {
    lineStreamer *lineStreamer
}


func NewAsciiNumberProvider(file string) *asciiNumberProvider{
    var anp = &asciiNumberProvider{}
    anp.lineStreamer = NewLineStreamer()
    anp.lineStreamer.Open(file)
    return anp
}

func (numProvider *asciiNumberProvider) GetNext() []string{
    
    //TODO: complete
    return nil
}


