package service

import (
	"github.com/pudongping/gin-blog-service/internal/dao"
	"github.com/pudongping/gin-blog-service/internal/model"
	"github.com/pudongping/gin-blog-service/internal/request"
	"github.com/pudongping/gin-blog-service/pkg/app"
)

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func (svc *Service) GetArticle(param *request.ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAID(article.Id)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagId, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}

	return &Article{
		ID:            article.Id,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           &tag,
	}, nil
}

func (svc *Service) GetArticleList(param *request.ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	articleCount, err := svc.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := svc.dao.GetArticleListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:            article.ArticleID,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{Id: article.TagID}, Name: article.TagName},
		})
	}

	return articleList, int(articleCount), nil
}

func (svc *Service) CreateArticle(param *request.CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.Id, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateArticle(param *request.UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		ModifiedBy:    param.ModifiedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteArticle(param *request.DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}
