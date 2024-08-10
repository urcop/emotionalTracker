package app

import (
	v1 "github.com/FoodMoodOTG/examplearch/api/v1"
	_ "github.com/FoodMoodOTG/examplearch/docs"
	"github.com/FoodMoodOTG/examplearch/services/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"log"
	"runtime"
	"strings"
	"time"
)

type HttpServer struct {
	app *fiber.App
}

type Server interface {
	Start()
}

func NewHttpServer() Server {
	app := fiber.New(fiber.Config{
		BodyLimit:         1024 * 1024 * 50,
		AppName:           "FoodMoodExample",
		StreamRequestBody: true,
	})

	var methods = []string{fiber.MethodGet, fiber.MethodPost, fiber.MethodPut, fiber.MethodDelete, fiber.MethodOptions}
	var headers = []string{fiber.HeaderAccept, fiber.HeaderAuthorization, fiber.HeaderContentType,
		fiber.HeaderContentLength, fiber.HeaderAcceptEncoding, "X-CSRF-Token"}

	corsConfig := cors.New(cors.Config{
		AllowOrigins: strings.Join([]string{"*"}, ", "),
		AllowMethods: strings.Join(methods, ", "),
		AllowHeaders: strings.Join(headers, ", "),
		MaxAge:       300,
	})

	app.Use(corsConfig)
	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:       "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Europe/Moscow",
		TimeInterval: 500 * time.Millisecond,
	}))

	return &HttpServer{app: app}
}

func (s *HttpServer) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	domainCtx := InitCtx().Make()

	cfg := config.Make()

	s.app.Use("", func(ctx *fiber.Ctx) error {
		ctx.Locals("context", domainCtx)
		return ctx.Next()
	})

	domainCtx.Services().Logger().Info("app context initialized", "op", "server.Start()")

	docs := s.app.Group("/docs")
	{
		docs.Get("/swagger/*", swagger.HandlerDefault)
	}
	domainCtx.Services().Logger().Info("swagger handler initialized", "op", "server.Start()")

	example := s.app.Group("/api/v1/example")
	{
		example.Post("/", v1.WrapHandler(v1.CreateExample))
		example.Get("/:id/", v1.WrapHandler(v1.GetExample))
		example.Get("/", v1.WrapHandler(v1.GetAllExamples))
	}

	domainCtx.Services().Logger().Info("auth handlers initialized", "op", "server.Start()")

	err := s.app.Listen(":" + cfg.HttpPort())
	if err != nil {
		log.Fatalf("failed to listen error due [%s]", err)
	}
}
