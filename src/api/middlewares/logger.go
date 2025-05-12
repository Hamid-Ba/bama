package middlewares

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/pkg/logging"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger, _ := logging.NewLogger(cfg.Logger)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger") {
			c.Next()
		} else {
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			start := time.Now() // start
			path := c.FullPath()
			raw := c.Request.URL.RawQuery

			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			c.Writer = blw
			c.Next()

			param := gin.LogFormatterParams{}
			param.TimeStamp = time.Now() // stop
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}
			param.Path = path

			keys := []logging.Field{}
			keys = append(keys, logging.Field{Key: "Path", Value: param.Path})
			keys = append(keys, logging.Field{Key: "ClientIp", Value: param.ClientIP})
			keys = append(keys, logging.Field{Key: "Method", Value: param.Method})
			keys = append(keys, logging.Field{Key: "Latency", Value: param.Latency})
			keys = append(keys, logging.Field{Key: "StatusCode", Value: param.StatusCode})
			keys = append(keys, logging.Field{Key: "ErrorMessage", Value: param.ErrorMessage})
			keys = append(keys, logging.Field{Key: "BodySize", Value: param.BodySize})
			keys = append(keys, logging.Field{Key: "RequestBody", Value: string(bodyBytes)})
			keys = append(keys, logging.Field{Key: "ResponseBody", Value: blw.body.String()})

			logger.Info("Requst-Response", keys...)
		}
	}
}
