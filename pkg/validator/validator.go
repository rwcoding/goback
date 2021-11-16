package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/rwcoding/goback/pkg/config"

	"log"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var translator ut.Translator

func Init() {
	var err error
	if config.GetLang() == "en" {
		uni := ut.New(en.New())
		translator, _ = uni.GetTranslator("en")
		validate = validator.New()
		err = en_translations.RegisterDefaultTranslations(validate, translator)
	} else {
		uni := ut.New(zh.New())
		translator, _ = uni.GetTranslator("zh")
		validate = validator.New()
		err = zh_translations.RegisterDefaultTranslations(validate, translator)
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Verify 验证一个结构体
func Verify(data interface{}) []string {
	if validate == nil {
		Init()
	}
	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return []string{err.Error()}
		}
		var e []string
		for _, err := range err.(validator.ValidationErrors) {
			e = append(e, err.Translate(translator))
		}
		return e
	}
	return nil
}

// VerifyField 验证单个变量
func VerifyField(field interface{}, tag string) string {
	if validate == nil {
		Init()
	}
	errs := validate.Var(field, tag)
	if errs != nil {
		errs := errs.(validator.ValidationErrors)
		ms := errs.Translate(translator)
		for _, v := range ms {
			return v
		}
	}
	return ""
}

// VerifyPost 验证POST请求表单
// rules 为验证规则，字段为key，规则为值，
// 例如：{"username": "required", "email": "required,email"}
func VerifyPost(c *gin.Context, rules map[string]string) []string {
	var ret []string
	for k, v := range rules {
		msg := VerifyField(c.PostForm(k), v)
		if msg != "" {
			ret = append(ret, k+msg)
		}
	}
	return ret
}
