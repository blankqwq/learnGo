package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translate "github.com/go-playground/validator/v10/translations/en"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"

	"net/http"
)

var trans ut.Translator

//清理结构体名称
func removeTopName(fieldErr map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fieldErr {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func InitTrans(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册tag翻译器
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("json")
			//name := strings.SplitN(field.Name, ".", 2)[0]
			if name == "-" {
				return ""
			}
			return strings.ToLower(name)
		})

		//创建翻译器
		zhT := zh.New()
		enT := en.New()
		// 创建通用翻译器:适配器
		uni := ut.New(enT, zhT, enT)
		// 获取翻译器
		trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", local)
		}
		// 将翻译器注册到 通用翻译器中
		switch local {
		case "en":
			return en_translate.RegisterDefaultTranslations(v, trans)
		case "zh":
			return zh_translate.RegisterDefaultTranslations(v, trans)
		default:
			return zh_translate.RegisterDefaultTranslations(v, trans)
		}
	}
	return
}

func main() {
	// 实例化对象
	err := InitTrans("zh")
	if err != nil {
		panic(err)
	}
	route := gin.Default()
	route.POST("/json", testJson)
	route.POST("/register", register)
	// 启动服务
	err = route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func register(c *gin.Context) {
	var register RegisterForm
	if err := c.ShouldBindJSON(&register); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			fmt.Printf("err ")
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"error": removeTopName(errs.Translate(trans)),
		})
		return
	}
	c.JSON(http.StatusOK, register)
}

type User struct {
	Username string `json:"username" binding:"required,min=3,max=10" `
	Password int    `json:"password" binding:"required" `
}

type RegisterForm struct {
	Username      string `json:"username" binding:"required,min=3,max=10" `
	Age           int32  `json:"age" binding:"required" `
	Password      string `json:"password" binding:"required" `
	CheckPassword string `json:"check_password" binding:"required,eqfield=Password" `
}

func testJson(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
