

package io

import (
    "fmt"
    "testing" 
)


func testBatch (lines []string, t *testing.T){
     for i, val := range lines{
            // fmt.Println(val)
            
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
    
    for  {
        lines = ls.ReadLines(4)
        
        if len(lines) < 4 {
            break
        }
        
        testBatch(lines, t)
    }
}


