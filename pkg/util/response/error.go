package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Meta
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error         `json:"-"`
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_FORBIDDEN            = "forbidden"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
	E_EXPIRED_LOGIN        = "force to login"
	E_CANCEL_ENTITY        = "cancel entity"
)

type errorHelper struct {
	Duplicate           Error
	NotFound            Error
	RouteNotFound       Error
	UnprocessableEntity Error
	Unauthorized        Error
	Forbidden           Error
	BadRequest          Error
	Validation          Error
	InternalServerError Error
	ExpiredLogin        Error
	CancelEntity        Error
}

var errorConstant errorHelper = errorHelper{
	Duplicate: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Status Conflict",
				Message: "",
			},
		},
		Code: http.StatusConflict,
	},
	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Data not found",
				Message: "",
			},
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Route not found",
				Message: "",
			},
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Invalid parameters or payload",
				Message: "",
			},
		},
		Code: http.StatusUnprocessableEntity,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Unauthorized, please login",
				Message: "",
			},
		},
		Code: http.StatusUnauthorized,
	},
	Forbidden: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Forbiden",
				Message: "",
			},
		},
		Code: http.StatusForbidden,
	},
	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Bad Request",
				Message: "",
			},
		},
		Code: http.StatusBadRequest,
	},
	Validation: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Invalid parameters or payload",
				Message: "",
			},
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  "Something bad happened",
				Message: "",
			},
		},
		Code: http.StatusInternalServerError,
	},
}

func ErrorBuilder(res Error, message error) *Error {
	res.ErrorMessage = message
	res.Response.Meta.Message = message.Error()
	return &res
}

func CustomErrorBuilder(code int, data any, message string, status string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  status,
				Message: message,
				Data:    data,
			},
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(Constant.Error.InternalServerError, err)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d, msg: %s", e.Code, e.Response.Meta.Message)
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) Send(c echo.Context) error {
	// type Payload struct {
	// 	Type string      `json:"type"`
	// 	Data any `json:"data"`
	// }
	// errs := []Payload{}
	// requestPayload :=
	logrus.Error(e.Response.Meta.Message)
	// body, err := ioutil.ReadAll(c.Request().Body)
	// if err != nil {
	// 	logrus.Warn("error read body, message: ", e.Error())
	// }

	// bHeader, err := json.Marshal(c.Request().Header)
	// if err != nil {
	// 	logrus.Warn("error read header, message: ", e.Error())
	// }

	// log.InsertErrorLog(c.Request().Context(), &log.LogError{
	// 	ID:           shortid.MustGenerate(),
	// 	Header:       string(bHeader),
	// 	Body:         string(body),
	// 	URL:          c.Request().URL.Path,
	// 	HttpMethod:   c.Request().Method,
	// 	ErrorMessage: e.ErrorMessage.Error(),
	// 	Level:        "Error",
	// 	AppName:      os.Getenv("APP"),
	// 	Version:      os.Getenv("VERSION"),
	// 	Env:          os.Getenv("ENV"),
	// 	CreatedAt:    time.Now().In(time.Local).Local().UTC(),
	// })

	return c.JSON(e.Code, e.Response)
}
