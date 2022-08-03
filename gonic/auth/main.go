package main
//$ curl -H 'Content-type:application/json' http://foo:bar@localhost:8080/admin -d '{"value":"abc1"}'
import (
	"net/http"
"fmt"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	p := fmt.Println
    pf := fmt.Printf
    p("")
    pf("")

	// Ping test
	r.POST("/ping", func(c *gin.Context) {
            ids := c.QueryMap("names")
            names := c.PostFormMap("ids")

            fmt.Printf("ids: %#v; names: %#v", ids, names)
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
        p(db)
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
    // curl -D-  -H 'Content-Type:application/json' 'http://foo:bar@127.0.0.1:8080/admin' -d '{"value":"1"}'
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
            Value string `json:"value" binding:"required"`
		}

        if err:=c.Bind(&json);err == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": user+"ok"})
		}else{
            p(err)
			c.JSON(http.StatusOK, gin.H{"status": "noset"})
        }
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
