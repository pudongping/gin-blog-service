package model

type ArticleTag struct {
	*Model
	// 文章 ID  int(11) is_nullable NO
	ArticleId uint32 `json:"article_id"`
	// 标签 ID  int(10) unsigned is_nullable NO
	TagId uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
