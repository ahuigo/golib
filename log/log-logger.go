package main
import (
    "bytes"
    "fmt"
    "log"
)
var buf bytes.Buffer
var logger = log.New(&buf, "prefix: ", log.Lshortfile)
func main() {

    logger.Print("log with file:lineno!")

    fmt.Print(&buf)
}
