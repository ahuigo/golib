package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func panicApi(c *gin.Context) {
	filepath := c.DefaultQuery("path", "tmp/a.txt")
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		err = errors.Wrap(err, "bad file argument!")
		panic(err)
		// c.String(http.StatusOK, err.Error())
	}
	res := string(buf)
	c.String(http.StatusOK, res)

}

type HTTPError struct {
	Error string
}

type User struct {
	ID      int    `json:"id" example:"1" format:"int64"`
	Name    string `json:"name" form:"name" example:"Alex"`
	Country string
	// time rfc3339
	Time time.Time `json:"time" form:"time"`
	// 这是city 说明
	City  string      `form:"city" example:"Beijing"`
	Citys []string    `form:"city" example:"Bj,Tj"`
	Extra interface{} `json:"extra"`
}

// Param Type: query path header body formData
// Refer to: https://github.com/swaggo/swag/blob/master/README.md#param-type

// @Summary      获取用户
// @Description  获取用户详情
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 Cookie header string  false "token"     default(token=xxx)
// @Param        id    path      int  true  "Account ID" Enums(1, 2, 3)
// @Param        name  query      int  true  "Account name" default("Alex")
// @Success      200  {object}  User
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {string}  string "500 error"
// @Router       /user/{id} [get]
func GetUser(c *gin.Context) {
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("bind error:", err)
	}
	var res User
	if user.ID == 1 {
		res = User{
			Name: "Alex",
		}
	} else {
		res = User{
			Name: "Other user",
		}
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 添加用户
// @Tags user
// @Param body-params body User true "Add user"
// @Success 200 {object} User
// @Failure 400  {object}  HTTPError
// @Router /user [post]
func AddUser(c *gin.Context) {
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("bind error:", err)
	}
	c.JSON(http.StatusOK, user)
}
