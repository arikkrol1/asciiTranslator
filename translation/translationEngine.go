
package translation

import(
    // "fmt"
    "github.com/arikkrol/asciiTranslator/translation/io"
)



type TranslationEngine struct{
    numberProvider *io.IInputProvider 
    outputHandler *io.IOutputHandler
}

