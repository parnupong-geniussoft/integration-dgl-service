package controllers

import (
	"integration-dgl-service/modules/dgl/entities"
	"integration-dgl-service/modules/dgl/usecases"

	"github.com/gofiber/fiber/v2"
)

type dglController struct {
	DglUsecase usecases.DglUsecase
}

func NewDglController(r fiber.Router, dglUsecase usecases.DglUsecase) {
	controllers := &dglController{
		DglUsecase: dglUsecase,
	}
	r.Post("/", controllers.DglMapper)
}

// @Summary Mapping dgl
// @Description Mapping dgl using the DGLRecord
// @Tags Dgl
// @Accept json
// @Produce json
// @Param body body entities.DGLRecord true "Dgl request object"
// @Success 200 {object} entities.DGLRecord "OK"
// @Router /dgl [post]
func (h *dglController) DglMapper(ctx *fiber.Ctx) error {
	body := new(entities.DGLRecord)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	body.MarketingCode = body.ProductMarketCode
	body.ProductMarketCode = ""

	result, err := h.DglUsecase.DglMapper(body)
	if err != nil {
		ctx.Context().SetStatusCode(fiber.StatusInternalServerError)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(result)
}
