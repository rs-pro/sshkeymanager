package api

import (
	"bytes"
	"io"
	"io/ioutil"

	"log"

	"github.com/gin-gonic/gin"
)

// Adapted from:
// https://stackoverflow.com/a/38548555/679778
// https://github.com/gin-gonic/gin/issues/961#issuecomment-557931409

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		log.Println("Request:", c.Request.Header, string(body))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		log.Println("Response:", c.Writer.Status(), blw.body.String())
	}
}
