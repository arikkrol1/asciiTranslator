

package io

import (
    "fmt"
    "testing" 
)


func testBatch (lines []string, t *testing.T){
     for i, val := range lines{
            //fmt.Println(val)
            
            if i == 3 {
                if val != "" {
                    fmt.Println("every 4th line should be blank")
                    t.FailNow()
                }
            } else {
                if val == "" {
                    fmt.Println("only 4th line should be blank")
                    t.FailNow()
                }
            }
        }
}

func TestStreaming(t *testing.T) {
    var ls = NewLineStreamer()
    
    ls.Open("./testdata/good.txt")
    
    var lines []string
    var moreLines = true
    
    for moreLines {
        lines, moreLines = ls.ReadLines(4)
        testBatch(lines, t)
    }
}




func TestGetNext(t *testing.T) {
    numProvider := NewAsciiNumberProvider("./testdata/good.txt")
    
    numRepresentation := numProvider.GetNext()
    
    for _,val := range numRepresentation{
        fmt.Println(val)
    }
    
    // t.FailNow()
}