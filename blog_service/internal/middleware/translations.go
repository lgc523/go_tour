package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := ctx.GetHeader("locale")
		translator, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, translator)
				break
			case "zh_Hant_TW":
				_ = tw_translations.RegisterDefaultTranslations(v, translator)
				break
			case "ja":
				_ = ja_translations.RegisterDefaultTranslations(v, translator)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, translator)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, translator)
				break
			}
			ctx.Set("trans", translator)
		}
		ctx.Next()
	}
}
