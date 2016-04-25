package io

import (
    "os"
    // "fmt"
    "bufio"
)


type lineStreamer struct {
	inFile *os.File
    scanner *bufio.Scanner
}

//NewLineStreamer ctor
func NewLineStreamer() *lineStreamer{
    return &lineStreamer{}
}

func (ls *lineStreamer) Open(file string){
    inFile, err := os.Open(file)
    if err != nil {
        panic("cant find or open input file " + file)
    }
    
    ls.inFile = inFile
    ls.scanner = bufio.NewScanner(ls.inFile)
}

func (ls *lineStreamer) Close(){
    ls.inFile.Close()
}

//ReadLines returns the lines read from the file + true if there are more lines to read or false otherwise 
func (ls *lineStreamer) ReadLines(linesNum int) []string {
    //length = 0, capacity = linesNum
    lines := make([]string, 0, linesNum)
    
	for  i := 0; i < linesNum; i++ {
        if ls.scanner.Scan() {
            lines =  append(lines, ls.scanner.Text())
        } else {
            break
        }
	}
    
    return lines
}





