package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_hexagonal_architecture/src/serializer/sum_serializer/json"
	"github.com/maei/golang_hexagonal_architecture/src/service/sum_service"
	"io/ioutil"
	"net/http"
)

type SumControllerInterface interface {
	NewCompute(c echo.Context) error
	FindResult(c echo.Context) error
}

type sumController struct {
	sumService sum_service.SumServiceInterface
}

func NewSumController(service sum_service.SumServiceInterface) SumControllerInterface {
	return &sumController{
		sumService: service,
	}
}

func (t *sumController) serializer() sum_service.SumSerializerInterface {
	return &json.Serial{}
}

func (t *sumController) NewCompute(c echo.Context) error {
	defer c.Request().Body.Close()

	jsonBody := c.Request().Body
	bs, err := ioutil.ReadAll(jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("error while converting to byte-slice"))
	}
	req, decErr := t.serializer().Decode(bs)
	if decErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("error while decoding to slice"))
	}
	saveErr := t.sumService.Compute(req)
	if saveErr != nil {
		return c.JSON(http.StatusInternalServerError, saveErr)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "compute save",
	})
}

func (t *sumController) FindResult(c echo.Context) error {
	code := c.Param("code")
	sumResp, err := t.sumService.FindResult(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]*sum_service.SumResponse{
		"message": sumResp,
	})
}
