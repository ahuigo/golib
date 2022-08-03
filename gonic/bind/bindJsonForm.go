package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "encoding/json"
    "time"

)
// 从 JSON 绑定
type Login struct {
    // 小写不能Bind
    //username string `form:"user" json:"user" binding:"required"`
    User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
    Arg  json.RawMessage `json:"arg"`
    BeginTime time.Time `form:"begin_time" json:"begin_time"`
    EndTime *time.Time `form:"end_time" json:"end_time"`
}

// $ curl   '0:8088/loginForm?user=manu&password=1234&start_time=&a=2021-11-15T00%3A00%3A00%2B08%3A00' -d '{}'
func main() {
    router := gin.Default()

    // 绑定 JSON 的示例 ({"user": "manu", "password": "123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        if err := c.ShouldBindJSON(&json); err == nil {
            fmt.Printf("arg:%T,%v\n", json.Arg, string(json.Arg))
            if json.User == "manu" && json.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // 一个 HTML 表单绑定的示例 (user=manu&password=123)
    router.POST("/loginForm", func(c *gin.Context) {
        var form Login
        // 这个将通过 content-type 头去推断绑定器使用哪个依赖: bindJSON/Form
        if err := c.ShouldBind(&form); err == nil {
            if form.User == "manu" && form.Password == "123" {
                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
            } else {
                fmt.Printf("req parms:%#v\n",form)
                fmt.Printf("req parms:%v\n",form.BeginTime)
                fmt.Printf("req parms:%v\n",form.EndTime)
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // 监听并服务于 0.0.0.0:8080
    router.Run(":8088")
}

