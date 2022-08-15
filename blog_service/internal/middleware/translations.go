package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
)

func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := ctx.GetHeader("locale")
		//translator, _ := uni.GetTranslator(locale)
		//v, ok := binding.Validator.Engine().(*validator.Validate)
		//if ok {
		//	switch locale {
		//	case "zh":
		//		_ = zh_translations.RegisterDefaultTranslations(v, translator)
		//		break
		//	case "zh_Hant_TW":
		//		_ = tw_translations.RegisterDefaultTranslations(v, translator)
		//		break
		//	case "ja":
		//		_ = ja_translations.RegisterDefaultTranslations(v, translator)
		//		break
		//	case "en":
		//		_ = en_translations.RegisterDefaultTranslations(v, translator)
		//		break
		//	default:
		//		_ = zh_translations.RegisterDefaultTranslations(v, translator)
		//		break
		//	}
		trans, found := global.Ut.GetTranslator(locale)
		if found {
			ctx.Set("trans", trans)
		} else {
			ctx.Set("trans", "en")
		}
		ctx.Next()
	}
}
