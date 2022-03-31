package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string  // 需要验证的字段
	Message string  // 验证失败之后的错误信息
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v) // 进行参数绑定和入参校验
	// 如果发生错误，则使用 Translations 中间件进行翻译错误消息体
	if err != nil {
		v := c.Value("trans")  // 从上下文中拿到翻译器
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)  // verrs 为验证器验证参数失败后的错误信息
		if !ok {
			return false, errs
		}

		// 将错误信息进行翻译
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
