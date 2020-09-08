package jwt

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// 用超时包装请求上下文
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		//匿名函数
		defer func() {
			// 检查是否达到上下文超时
			if ctx.Err() == context.DeadlineExceeded {

				// 写入响应并中止请求
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": e.QUESTTIMEOUT,
					"msg":  e.GetMsg(e.QUESTTIMEOUT),
					"data": nil,
				})
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			//完成后取消以清除资源
			cancel()
		}()

		// 用上下文包装的请求替换请求
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// get the underlying request context
		ctx := c.Request.Context()

		// create the response data type to use as a channel type
		type responseData struct {
			status int
			body   map[string]interface{}
		}

		// create a done channel to tell the request it's done
		doneChan := make(chan responseData)

		// here you put the actual work needed for the request
		// and then send the doneChan with the status and body
		// to finish the request by writing the response
		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()

		// non-blocking select on two channels see if the request
		// times out or finishes
		select {

		// if the context is done it timed out or was cancelled
		// so don't return anything
		case <-ctx.Done():
			return

			// if the request finished then finish the request by
			// writing the response
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}
