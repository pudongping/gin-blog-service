package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/pudongping/gin-blog-service/pkg/tracer"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/internal/model"
	"github.com/pudongping/gin-blog-service/internal/routers"
	"github.com/pudongping/gin-blog-service/pkg/logger"
	"github.com/pudongping/gin-blog-service/pkg/setting"
)

var (
	port         string
	runMode      string
	config       string
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	// 初始化加载配置信息
	err = setupSetting()
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

}

// @title gin-blog-service 博客系统
// @version 1.0
// @description gin-blog-service 学习 gin 写的一个博客系统
// @termsOfService https://github.com/pudongping/gin-blog-service
func main() {

	// 当携带 -version 参数执行二进制文件时，打印版本信息
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}

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

	// 优雅的重启和停止
	gracefulShutdown(s)

}

func gracefulShutdown(srv *http.Server) {
	// 优雅的重启和停止
	// see gin web framework document examples : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal)
	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	// kill 不加参数发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，因此不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 最大时间控制，用于通知该服务端它有 5 秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

// setupSetting 加载配置文件
func setupSetting() error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...) // 加载配置文件
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
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

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
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

// setupTracer 链路追踪
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"gin-blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
