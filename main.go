package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/internal/model"
	"github.com/pudongping/gin-blog-service/internal/routers"
	"github.com/pudongping/gin-blog-service/pkg/setting"
)

func init() {
	// 初始化加载配置信息
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// 初始化连接数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
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

	// 将读取到的配置信息绑定到对应的结构体中
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

// setupDBEngine 初始化设置数据库连接
func setupDBEngine() error {
	var err error
	// 这里需要注意：不能写成 ==> global.DBEngine, err := model.NewDBEngine(global.DatabaseSetting)
	// 因为 `:=` 会重新声明并创建了左侧的新局部变量，因此在其它包中调用 global.DBEngine 变量时，它仍然是 nil
	// 因为根本就没有赋值到包全局变量 global.DBEngine 上
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
