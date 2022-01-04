package model

type Article struct {
	*Model
	// 文章标题  varchar(100) is_nullable YES
	Title string `json:"title"`
	// 文章简述  varchar(255) is_nullable YES
	Desc string `json:"desc"`
	// 封面图片地址  varchar(255) is_nullable YES
	CoverImageUrl string `json:"cover_image_url"`
	// 文章内容  longtext is_nullable YES
	Content string `json:"content"`
	// 状态 0 为禁用、1 为启用  tinyint(3) unsigned is_nullable YES
	State uint8 `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
