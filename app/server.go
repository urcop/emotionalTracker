package app

import (
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	v1 "github.com/urcop/emotionalTracker/api/v1"
	_ "github.com/urcop/emotionalTracker/docs"
	"github.com/urcop/emotionalTracker/services/config"
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
		AppName:           "EmotionalTracker",
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

	user := s.app.Group("/api/v1/user")
	{
		user.Post("/", v1.WrapHandler(v1.CreateUser))
		user.Get("/telegram/", v1.WrapHandler(v1.GetUserByTelegramId))
		user.Get("/:id/zodiac/", v1.WrapHandler(v1.GetUserZodiacSign))
		user.Put("/:id/birthday", v1.WrapHandler(v1.UpdateUserBirthday))
		user.Get("/:telegram_id/horoscope/", v1.WrapHandler(v1.GetUserHoroscope))
		user.Get("/:id/", v1.WrapHandler(v1.GetUser))
		user.Get("/", v1.WrapHandler(v1.GetAllUsers))
		user.Put("/", v1.WrapHandler(v1.UpdateUser))
		user.Delete("/:id/", v1.WrapHandler(v1.DeleteUser))
	}

	zodiac := s.app.Group("/api/v1/zodiac")
	{
		zodiac.Get("/", v1.WrapHandler(v1.GetZodiacSign))
		zodiac.Get("/horoscope", v1.WrapHandler(v1.GetHoroscope))
	}

	domainCtx.Services().Logger().Info("api handlers initialized", "op", "server.Start()")

	err := s.app.Listen(":" + cfg.HttpPort())
	if err != nil {
		log.Fatalf("failed to listen error due [%s]", err)
	}
}
