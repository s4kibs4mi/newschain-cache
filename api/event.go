package api

import (
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/newschain-cache/app"
	"github.com/s4kibs4mi/newschain-cache/data"
	"github.com/s4kibs4mi/newschain-cache/ent"
	"net/http"
	"strconv"
	"strings"
)

func listEvents(ctx echo.Context) error {
	resp := response{}
	var err error

	limit := 9
	currentPageP := ctx.QueryParam("page")
	currentPage, err := strconv.ParseInt(currentPageP, 32, 10)
	if err != nil || currentPage <= 0 {
		currentPage = 1
	}

	from := (int(currentPage) * limit) - limit

	db := app.DB()

	var posts []*ent.Post

	query := strings.TrimSpace(ctx.QueryParam("query"))
	if query == "" {
		posts, err = data.ListEvents(db.Post, from, limit)
		if err != nil {
			return resp.ServerErrorAsJSON(ctx, "Database query failed", http.StatusInternalServerError, err)
		}
	} else {
		posts, err = data.SearchEvents(db.Post, query, from, limit)
		if err != nil {
			return resp.ServerErrorAsJSON(ctx, "Database query failed", http.StatusInternalServerError, err)
		}
	}

	resp.Status = http.StatusOK
	resp.Data = posts
	return resp.ServerJSON(ctx)
}
