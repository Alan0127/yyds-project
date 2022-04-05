package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Test(ctx *gin.Context) {
	//a, _ := ctx.Get("myProject_ctx_key")

	ctx.Set("returnCode", 9999)
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, "0000")
}
