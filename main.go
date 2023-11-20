package main

import (
	"net/http"

	"github.com/HiWay-Media/hwm-go-utils/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	zapLogger := log.GetLogger("debug")
	//
	f := fiber.New()
	f.Use(logger.New())
	f.Use(recover.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",
		AllowMethods: "GET, HEAD, OPTIONS, PUT, PATCH, POST, DELETE",
	}))
	v1 := f.Group("/api/v1.0/")
	setApi(f)
	setHealth(f)
	err := f.Listen("0.0.0.0:8000")
	zapLogger.Error(err.Error())
	//
}

func setHealth(f *fiber.App) {
	f.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	f.Head("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}
//
func setApi(f *fiber.App) {
	f.Get("/panic", func(c *fiber.Ctx) error {
		panic("crash to test")
	})
}
//