package dto

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err error `json:"-"` // low-level runtime error

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging

	HTTPStatusCode int   `json:"-"`              // http response status code
	AppCode        int64 `json:"code,omitempty"` // application-specific error code
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
	}
}

type OkResponse struct {
	StatusText     string `json:"status"`
	Message        string `json:"message"`
	AppCode        int64  `json:"code,omitempty"`
	HTTPStatusCode int    `json:"-"`
}

func (e *OkResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func OkResponseData(msg string) render.Renderer {
	return &OkResponse{
		HTTPStatusCode: http.StatusOK,
		StatusText:     "Ok",
		Message:        msg,
	}
}
