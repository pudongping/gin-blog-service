package model

import (
	"github.com/pudongping/gin-blog-service/pkg/app"
)

type Tag struct {
	*Model
	// 标签名称  varchar(100) is_nullable YES
	Name string `json:"name"`
	// 状态 0 为禁用、1 为启用  tinyint(3) unsigned is_nullable YES
	State uint8 `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
