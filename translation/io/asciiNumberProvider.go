package io

import (
    // "os"
    // "fmt"
    // "bufio"
)

const linesToRead = 4
const digitHeight = 3
const digitWidth = 3
const digits = 9

type asciiNumberProvider struct {
    lineStreamer *lineStreamer
}

//NewAsciiNumberProvider ctor
func NewAsciiNumberProvider(file string) *asciiNumberProvider{
    var anp = &asciiNumberProvider{}
    anp.lineStreamer = NewLineStreamer()
    anp.lineStreamer.Open(file)
    return anp
}

func (numProvider *asciiNumberProvider) GetNext() []string{
    lines := numProvider.lineStreamer.ReadLines(linesToRead)
    
    if len(lines) < linesToRead {
        return nil
    }
    
    res := make([]string, digits)
    for i := 0;  i < digits; i++ {
        res[i] = getDigit(lines, i * digitWidth)
    }
    return res
}

func (numProvider *asciiNumberProvider) Close(){
    numProvider.lineStreamer.Close()
}

func getDigit(lines []string, index int) string{
    digitArr := make([]byte, digitWidth * digitHeight)
    arrPos := 0
    for row := 0;  row < digitHeight; row++ {
        for col := index; col < index + digitWidth; col++ {
            char := lines[row][col]
            digitArr[arrPos] = char
            arrPos++
        }
    }
    
    res := string(digitArr)
    return res
}


