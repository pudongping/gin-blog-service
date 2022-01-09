package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/pudongping/gin-blog-service/docs"
	"github.com/pudongping/gin-blog-service/internal/middleware"

	v1 "github.com/pudongping/gin-blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations()) // 国际化处理中间件

	// swaggerUrl := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)  // 创建标签
		apiv1.DELETE("/tags/:id", tag.Delete)  // 删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)  // 更新指定标签
		apiv1.GET("/tags", tag.List)  // 获取标签列表

		apiv1.POST("/articles", article.Create)  // 创建文章
		apiv1.DELETE("/articles/:id", article.Delete)  // 删除指定文章
		apiv1.PUT("/articles/:id", article.Update)  // 更新指定文章
		apiv1.GET("/articles/:id", article.Get)  // 获取指定文章
		apiv1.GET("/articles", article.List)  // 获取文章列表
	}

	return r
}
