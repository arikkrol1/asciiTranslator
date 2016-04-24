package io

import (
    "os"
    "fmt"
    "bufio"
)


type lineStreamer struct {
	inFile *os.File
    scanner *bufio.Scanner
}

func NewLineStreamer() *lineStreamer{
    return &lineStreamer{}
}

func (ls *lineStreamer) Open(file string){
    inFile, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    ls.inFile = inFile
    ls.scanner = bufio.NewScanner(ls.inFile)
}

func (ls *lineStreamer) Close(){
    ls.inFile.Close()
}

//ReadLines returns the lines read from the file + true if there are more lines to read or false otherwise 
func (ls *lineStreamer) ReadLines(linesNum int) ([]string, bool) {
    lines := make([]string, linesNum)
    
	for  i := 0; i < linesNum; i++ {
        if ls.scanner.Scan() {
            lines[i] = ls.scanner.Text() 
        } else {
            ls.Close()
            return lines, false
        }
	}
    
    return lines, true
}







