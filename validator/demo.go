package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func multiLangBindingHandler(c *gin.Context) {

	var person Person

	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := validate.Struct(person); err != nil {

		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			// can translate each error one at a time.
			sliceErrs = append(sliceErrs, e.Translate(tans))
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   errs.Translate(tans),
			"sliceErrs": sliceErrs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"personInfo": person,
	})
}
func langMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		local := c.DefaultQuery("lang", "en")
		log.Println("lang:", local)
		tans, _ := uni.GetTranslator(local)
		switch local {
		case "en":
			en_translations.RegisterDefaultTranslations(validate, tans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(validate, tans)
		default:
			en_translations.RegisterDefaultTranslations(validate, tans)
		}

		c.Next()
	}
}

func main() {
	zh := zh2.New()
	en := en2.New()
	uni = ut.New(en, zh)

	validate = validator.New()
	router := gin.Default()
	router.Use(langMiddle())
	router.GET("/testMultiLangBinding", multiLangBindingHandler)
	router.Run(":9999")
}
