package main

import (
	"hello/algoritma"
	"hello/basic_coding/database"
	"hello/basic_coding/router"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	database.Init()
	e := router.Init()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/fibo/:n", func(c echo.Context) error {
		n, err := strconv.Atoi(c.Param("n"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		fibs := algoritma.Fibonaci(n)
		response := map[string]interface{}{
			"fibs": fibs,
		}
		return c.JSONPretty(http.StatusOK, response, "	")
	})

	e.GET("/prime/:n", func(c echo.Context) error {
		n, err := strconv.Atoi(c.Param("n"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		primes := algoritma.GetPrimesNumber(n)
		response := map[string]interface{}{
			"primes": primes,
		}
		return c.JSONPretty(http.StatusOK, response, "	")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
