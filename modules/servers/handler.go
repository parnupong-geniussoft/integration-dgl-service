package servers

import (
	_dglControllers "integration-dgl-service/modules/dgl/controllers"
	_dglRepositories "integration-dgl-service/modules/dgl/repositories"
	_dglUsecases "integration-dgl-service/modules/dgl/usecases"
	"integration-dgl-service/modules/middlewares"

	_ "integration-dgl-service/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (s *Server) MapHandlers() error {
	// Swagger UI
	s.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	s.App.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "ok"})
	})

	s.App.Use(middlewares.SystemLoggerMiddleware(*s.Log))
	s.App.Use(middlewares.DbLoggerMiddleware(*s.Log))
	s.App.Use(middlewares.RecoverMiddleware())
	s.App.Use(middlewares.BasicAuthMiddleware(s.Cfg.Auth.BasicAuthUsername, s.Cfg.Auth.BasicAuthPassword))

	// Public routes
	publicGroup := s.App.Group("/dgl")

	// Dgl Controller
	dglRepository := _dglRepositories.NewDglRepository(s.Db, s.C)
	dglUsecase := _dglUsecases.NewDglUsecase(dglRepository)
	_dglControllers.NewDglController(publicGroup, dglUsecase)

	// End point not found response
	s.App.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
