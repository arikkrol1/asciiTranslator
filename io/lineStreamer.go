package asciiTranslator

import (
    "os"
    "fmt"
    "bufio"
    // "regexp"
    // "strings"
    // "io/ioutil"
    // "github.com/revel/revel"
    // "encoding/json"
    //"bjjMap/app/models"
    // "bjjMap/app/dal"
    
    //"bjjMap/app/log"
)


type LineStreamer struct {
	inFile *os.File
    scanner *bufio.Scanner
}

func (ls LineStreamer) Open(file string){
    inFile, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    ls.inFile = inFile
    ls.scanner = bufio.NewScanner(ls.inFile)
    ls.scanner.Split(bufio.ScanLines)
}

func (ls LineStreamer) Close(){
    ls.inFile.Close()
}

//ReadLines returns the lines read from the file + true if there are more lines to read or false otherwise 
func (ls LineStreamer) ReadLines(linesNum int) ([]string, bool) {
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







