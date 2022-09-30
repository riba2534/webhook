package handler

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, _ := ioutil.ReadAll(c.Request.Body)
	c.String(200, string(data))
}
