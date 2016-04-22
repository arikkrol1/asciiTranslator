

package io

import (
    "fmt"
    "testing" 
)

func TestTestFunc(t *testing.T) {
    var ls = &LineStreamer{}
    
    
    
    
    ls.Open("./testdata/good.txt")
    
    //TODO: for some reason ls.scanner is nil here
    //fmt.Println(ls.scanner)
    // fmt.Println(ls.inFile)
    
    lines, _ := ls.ReadLines(3)
    
    for _, val := range lines{
        fmt.Println(val)
    }
    
}