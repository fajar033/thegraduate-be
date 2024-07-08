package helper

import (
	"thegraduate-server/model"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {

	report, ok := err.(*echo.HTTPError)

	if ok {
		c.JSON(report.Code, model.ResponseModelFailed{Message: report.Message, Status: "failed"})
		return
	}

	switch e := err.(type) {
	case *model.NotFoundError:
		c.JSON(404, model.ResponseModelFailed{
			Message: e.Error(),
			Status:  "NOT_FOUND",
		})
	case *model.ConflictError:
		c.JSON(409, model.ResponseModelFailed{
			Message: e.Message,
			Status:  "CONFLICT_ERROR",
		})
	case *model.BadRequestError:
		c.JSON(400, model.ResponseModelFailed{Message: e.Message, Status: "BAD_REQUEST"})
	case *model.ValidationError:

		c.JSON(400, model.ResponseModelFailed{
			Message: e.ErrMessage,
			Status:  "BAD_REQUEST",
		})
	default:

		c.JSON(500, model.ResponseModelFailed{
		
			Message: err.Error(),
			Status:  "ERROR",
		})
	}

}
