package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		// 通过 GetHeader 方法去获取约定的 header 参数 locale，用于判别当前请求的语言类别是 en 又或是 zh
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				// 调用 RegisterDefaultTranslations 方法将验证器和对应语言类型的 Translator 注册进来，实现验证器的多语言支持
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 将 Translator 存储到全局上下文中，便于后续翻译时的使用
			c.Set("trans", trans)
		}

		c.Next()
	}
}
