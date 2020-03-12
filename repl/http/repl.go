package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/lucasew/golisp/toolchain/default"
	"sync"
)

func main() {
	var err error
	tc := tdefault.NewDefaultToolchain(nil)
	var mutex sync.Mutex
	r := gin.Default()
	r.GET("/eval", func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			handleError(c, err)
			return
		}
		ast, err := tc.ParseBytes(buf)
		if err != nil {
			handleError(c, err)
			return
		}
		mutex.Lock()
		res, err := tc.Eval(ast)
		mutex.Unlock()
		c.JSON(200, gin.H{
			"result": res.Repr(),
		})
	})
	r.GET("/eval-spew", func(c *gin.Context) {
		buf := make([]byte, 4096)
		_, err = c.Request.Body.Read(buf)
		if err != nil {
			handleError(c, err)
			return
		}
		ast, err := tc.ParseBytes(buf)
		if err != nil {
			handleError(c, err)
			return
		}
		mutex.Lock()
		res, err := tc.Eval(ast)
		mutex.Unlock()
		c.JSON(200, gin.H{
			"result": spew.Sdump(res),
		})
	})
	r.Run()
}

func handleError(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}
