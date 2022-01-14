package service

import (
	"errors"

	"github.com/pudongping/gin-blog-service/internal/request"
)

// CheckAuth 授权检查
func (svc *Service) CheckAuth(param *request.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	// 存在用户
	if auth.Id > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
