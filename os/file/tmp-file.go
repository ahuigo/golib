package main
import (
    "io/ioutil"
    "os"
    //"path/filepath"
)


func createTempFile(dir string) (file *os.File, err error){
    //dir := filepath.Split(path)
    os.MkdirAll(dir, os.ModePerm)
    file, err = ioutil.TempFile(dir, "tmp_")
    return
    //defer os.Remove(file.Name())
}


func main(){
    os.Mkdir("tmp", 0700)
    path := "tmp/a/b/c/d"
    file, err := createTempFile(path)
    if err!=nil{
        panic(err)
    }
    println("Succssfully create file: ", file.Name())

}
