package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidErr struct {
	Key string
	Msg string
}

type ValidErrs []*ValidErr

func (v *ValidErr) Error() string {
	return v.Msg
}

func (v ValidErrs) Errors() []string {
	var errs []string
	for _, v := range v {
		errs = append(errs, v.Msg)
	}
	return errs
}

func (v ValidErrs) Error() string {
	return strings.Join(v.Errors(), ",")
}

func BindAndValid(ctx *gin.Context, v any) (bool, ValidErrs) {
	var errs ValidErrs
	err := ctx.ShouldBind(v)
	if err != nil {
		v := ctx.Value("trans")
		trans, _ := v.(ut.Translator)
		verse, ok := err.(validator.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verse.Translate(trans) {
			errs = append(errs, &ValidErr{
				key,
				value,
			})
		}
		return true, errs
	}
	return false, nil
}
