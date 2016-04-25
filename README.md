# klarna-interview-go


1. Install Go runtime from https://golang.org/dl/

   You should have Go version 1.4+
   
   Default location for linux is /usr/local/go
   
2. Make sure you have your GOPATH environment variable set to the installtion directory (/usr/local/go by default)  

   Also make sure that GOPATH/bin is in the system path so "go" command can be executed from anywhere. 

3. Execute "go get github.com/arikkrol1/asciiTranslator" from the commandline

   This will add the code to GOPATH/src/github.com/arikkrol1/asciiTranslator 

4. Change input and output file paths in app.conf 

5. All tests can be launched with "go test ./..." from asciiTranslator directory

6. Run the code using "go run main.go"

   You can also build an executable using "go build main.go" and run the executable 

7. This was tested on Windows 10 and Linux 14.04
