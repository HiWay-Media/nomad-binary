package main

import (

	"github.com/HiWay-Media/hwm-go-utils/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

)

func main() {
	zapLogger := log.GetLogger("debug")
	if err != nil {
		fmt.Printf("error initializing the logger %v\n", err)
	}
	sugaredLogger := zapLogger.Sugar()
	//
	f := fiber.New()
	f.Use(logger.New())
	f.Use(recover.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",
		AllowMethods: "GET, HEAD, OPTIONS, PUT, PATCH, POST, DELETE",
	}))
	//v1 := f.Group("/api/v1.0/")
	f.Head("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	err = f.Listen("0.0.0.0:8000")
	//
	
}