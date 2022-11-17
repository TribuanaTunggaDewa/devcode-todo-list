package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type successHelper struct {
	OK Success
}

var successConstant successHelper = successHelper{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Status:  "Request successfully proceed",
				Message: "",
				Data:    nil,
			},
		},
		Code: http.StatusOK,
	},
}

type successResponse struct {
	Meta
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

func SuccessBuilder(res Success, data any) *Success {
	res.Response.Meta.Data = data
	return &res
}

func CustomSuccessBuilder(code int, data any, message string, status string) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Status:  "success",
				Message: "success",
				Data:    data,
			},
		},
		Code: code,
	}
}

func SuccessResponse(data any) *Success {
	return SuccessBuilder(Constant.Success.OK, data)
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}
