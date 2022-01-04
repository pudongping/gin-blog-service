package main

import (
	"net/http"
	"time"

	"github.com/pudongping/gin-blog-service/internal/routers"
)

func main() {

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second, // 允许读取的最大时间
		WriteTimeout:   10 * time.Second, // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,          // 请求头的最大字节数
	}
	s.ListenAndServe()
}
