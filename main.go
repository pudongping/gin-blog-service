package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/internal/routers"
	"github.com/pudongping/gin-blog-service/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {

	gin.SetMode(global.ServerSetting.RunMode) // 设置 gin 的运行模式

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,  // 允许读取的最大时间
		WriteTimeout:   global.ServerSetting.WriteTimeout, // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                           // 请求头的最大字节数
	}

	s.ListenAndServe()
}

// setupSetting 加载配置文件
func setupSetting() error {
	newSetting, err := setting.NewSetting() // 加载配置文件
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
