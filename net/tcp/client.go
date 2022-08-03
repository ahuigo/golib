package main

import (
   "bufio"
   "fmt"
   "net"
   "os"
   "strings"
   "bytes"
   "io"
)

func main() {
   conn,err := net.Dial("tcp","127.0.0.1:9090")
   if err != nil {
      fmt.Println("Connect to TCP server failed ,err:",err)
      return
   }

   // 读取命令行输入
   inputReader := bufio.NewReader(os.Stdin)
   for {
      input,err := inputReader.ReadString('\n')
      if err != nil {
         fmt.Println("Read from console failed,err:",err)
         return
      }

      // 读取到字符"Q"退出
      str := strings.TrimSpace(input)
      if str == "Q"{
         break
      }

      // 响应服务端信息
      _,err = conn.Write([]byte(input))
      if err != nil{
         fmt.Println("Write failed,err:",err)
         break
      }
    var b *bytes.Buffer = bytes.NewBuffer([]byte("copy end"))
    println(b, io.Copy)
    _, err = io.Copy(conn, b) // copy b to stdout
   }
   
}
