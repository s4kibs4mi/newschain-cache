package api

import "github.com/labstack/echo/v4"

type response struct {
	Status int         `json:"-"`
	Title  string      `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

func (r *response) ServerJSON(ctx echo.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json")

	if err := ctx.JSON(r.Status, r); err != nil {
		return err
	}
	return nil
}

func (r *response) ServerErrorAsJSON(ctx echo.Context, title string, status int, err error) error {
	ctx.Response().Header().Set("Content-Type", "application/json")

	r.Title = title
	r.Errors = err
	r.Status = status
	if err := ctx.JSON(r.Status, r); err != nil {
		return err
	}
	return nil
}
