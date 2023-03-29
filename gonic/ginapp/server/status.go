package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func StatusServer(ctx *gin.Context) {
	code, _ := strconv.Atoi(ctx.Param("code"))
	ctx.Status(code)
}
