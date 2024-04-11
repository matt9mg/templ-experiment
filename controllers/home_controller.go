package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/matt9mg/templ-experiment/types"
	views_home "github.com/matt9mg/templ-experiment/views/home"
	"log"
)

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (*Home) Index(ctx fiber.Ctx) error {
	handler := adaptor.HTTPHandler(templ.Handler(views_home.Index("home")))
	return handler(ctx)
}

func (*Home) Save(ctx fiber.Ctx) error {
	var t *types.Test

	if err := ctx.Bind().JSON(&t); err != nil {
		log.Println(err)
	}

	handler := adaptor.HTTPHandler(templ.Handler(views_home.Save(t)))

	return handler(ctx)
}
