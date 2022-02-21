package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now()
		c.Next()
		endTime := time.Now()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(), // 当前的请求参数
			"url": c.Request.URL,
			"userAgent": c.Request.UserAgent(),
			"proto": c.Request.Proto,
			"header": c.Request.Header,
			"host": c.Request.Host,
			"remoteAddr": c.Request.RemoteAddr,
			"requestUri": c.Request.RequestURI,
			"response": bodyWriter.body.String(),    // 当前的请求结果响应主体
		}
		s := "access log: method: %s, status_code: %d, " +
			"begin_time: %d, end_time: %d, begin_time_date: %s, end_time_date: %s, code_execute_time: %s"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,                        // 当前的调用方法
			bodyWriter.Status(),                     // 当前的响应结果状态码
			beginTime.Unix(),                               // 调用方法的开始时间
			endTime.Unix(),                                 // 调用方法的结束时间
			beginTime.Format("2006-01-02 15:04:05"), // 格式化处理的开始时间
			endTime.Format("2006-01-02 15:04:05"),   // 格式化处理的结束时间
			time.Since(beginTime),                   // 代码执行时间
		)

	}
}
