package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	views_home "github.com/matt9mg/templ-experiment/views/home"
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
