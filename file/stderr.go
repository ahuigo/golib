import "fmt"
import "os"

func main() {
    str := "Message"
    fmt.Fprintln(os.Stderr, str)
    io.WriteString(os.Stderr, str)
    io.Copy(os.Stderr, bytes.NewBufferString(str))
    os.Stderr.Write([]byte(str))
}
