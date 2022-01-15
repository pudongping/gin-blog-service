package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pudongping/gin-blog-service/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/internal/routers/api"

	_ "github.com/pudongping/gin-blog-service/docs"
	"github.com/pudongping/gin-blog-service/internal/middleware"

	v1 "github.com/pudongping/gin-blog-service/internal/routers/api/v1"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog()) // 访问日志记录
		r.Use(middleware.Recovery())  // 程序异常处理
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout)) // 接口请求超时设置
	r.Use(middleware.Translations())                                          // 国际化处理中间件

	// swaggerUrl := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()

	r.POST("/upload/file", upload.UploadFile)                         // 上传文件接口
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) // 设置文件服务去提供静态资源的访问

	r.POST("/auth", api.GetAuth) // 鉴权接口

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())

	{
		apiv1.POST("/tags", tag.Create)       // 创建标签
		apiv1.DELETE("/tags/:id", tag.Delete) // 删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)    // 更新指定标签
		apiv1.GET("/tags", tag.List)          // 获取标签列表

		apiv1.POST("/articles", article.Create)       // 创建文章
		apiv1.DELETE("/articles/:id", article.Delete) // 删除指定文章
		apiv1.PUT("/articles/:id", article.Update)    // 更新指定文章
		apiv1.GET("/articles/:id", article.Get)       // 获取指定文章
		apiv1.GET("/articles", article.List)          // 获取文章列表
	}

	return r
}
