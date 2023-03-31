package commonginmiddleware

import (
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

const dictionary = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-+"

// RequestID is a middleware that generates a unique request ID for each request.
// The request ID is added to the request context and is also added to the
// response header.
func UseRequestID(num int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b strings.Builder
		for i := 0; i < num; i++ {
			b.WriteByte(dictionary[rand.Intn(len(dictionary))])
		}
		c.Set("requestId", b.String())
		c.Writer.Header().Set("X-Request-Id", b.String())
		c.Next()
	}
}
