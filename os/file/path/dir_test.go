package path
import (
	"os"
	"path"
	"runtime"
    "fmt"
	"io"
	"log"
	"path/filepath"
)

func TestCurrentDir1(){
    _, filename, _, _ := runtime.Caller(0)
    dir:=path.Dir(filename)
    println(dir)
}

func TestCurrentDir2(){
    exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
    println(dir)
}
