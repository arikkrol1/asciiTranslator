package io

import (
    // "os"
    // "fmt"
    // "bufio"
)


type asciiNumberProvider struct {
    lineStreamer *lineStreamer
    
    linesToRead int
    digitHeight int
    digitWidth int
    digits int
}

//NewAsciiNumberProvider ctor
func NewAsciiNumberProvider(config map[string]interface{}) *asciiNumberProvider{
    var numProvider = &asciiNumberProvider{}
    numProvider.lineStreamer = NewLineStreamer()
    numProvider.lineStreamer.Open(config["inputFile"].(string))
    
    numProvider.linesToRead = int(config["linesToRead"].(float64))
    numProvider.digitHeight = int(config["digitHeight"].(float64))
    numProvider.digitWidth = int(config["digitWidth"].(float64))
    numProvider.digits = int(config["digits"].(float64))
    
    return numProvider
}

func (numProvider *asciiNumberProvider) GetNext() []string{
    lines := numProvider.lineStreamer.ReadLines(numProvider.linesToRead)
    
    if len(lines) < numProvider.linesToRead {
        return nil
    }
    
    res := make([]string, numProvider.digits)
    for i := 0;  i < numProvider.digits; i++ {
        res[i] = numProvider.getDigit(lines, i * numProvider.digitWidth)
    }
    return res
}

func (numProvider *asciiNumberProvider) Close(){
    numProvider.lineStreamer.Close()
}

func (numProvider *asciiNumberProvider) getDigit(lines []string, index int) string{
    digitArr := make([]byte, numProvider.digitWidth * numProvider.digitHeight)
    arrPos := 0
    for row := 0;  row < numProvider.digitHeight; row++ {
        for col := index; col < index + numProvider.digitWidth; col++ {
            char := lines[row][col]
            digitArr[arrPos] = char
            arrPos++
        }
    }
    
    res := string(digitArr)
    return res
}


