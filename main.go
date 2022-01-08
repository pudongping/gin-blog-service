package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/internal/model"
	"github.com/pudongping/gin-blog-service/internal/routers"
	"github.com/pudongping/gin-blog-service/pkg/logger"
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

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

func main() {

	fmt.Printf("App server is running at: http://127.0.0.1:%s \n", global.ServerSetting.HttpPort)

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

// setupLogger 初始化日志系统
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName, // 日志文件名
		MaxSize:   600,      // 设置日志文件所允许的最大占用空间为 600MB
		MaxAge:    10,       // 日志文件最大生存周期为 10 天
		LocalTime: true,     // 设置日志文件名的时间格式为本地时间
	}, "", log.LstdFlags)

	return nil
}
