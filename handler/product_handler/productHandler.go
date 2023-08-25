package product_handler

import (
	"github.com/labstack/echo/v4"
	"web-service-test/entities"
	"web-service-test/usecases/product_usecase"
)

type ProductHandler struct {
	productUseCase product_usecase.ProductUseCase
}

func NewProductHandler(productUseCase product_usecase.ProductUseCase) ProductHandler {
	return ProductHandler{productUseCase: productUseCase}
}

func (ph ProductHandler) GetAllProducts(c echo.Context) error {
	ok, err, data := ph.productUseCase.GetAllProduct()
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "product not found",
			Result:  err,
		})
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "product found",
		Result:  data,
	})
}

func (ph ProductHandler) GetProductDetail(c echo.Context) error {
	paramProductId := c.QueryParam("id")
	ok, err, data := ph.productUseCase.GetProductDetail(paramProductId)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "product not found",
			Result:  err,
		})
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "product found",
		Result:  data,
	})
}
