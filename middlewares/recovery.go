package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/newschain-cache/log"
	"runtime"
)

func Recovery() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					log.Log().Errorln(err)
					stack := make([]byte, 10)
					length := runtime.Stack(stack, false)
					log.Log().Printf("[PANIC RECOVER] %v %s\n", err, stack[:length])
				}
			}()

			return next(c)
		}
	}
}
