package echomiddleware

import "github.com/labstack/echo"

func nopMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error { return next(c) }
}