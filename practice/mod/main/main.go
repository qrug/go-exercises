package main

import (
	"fmt"

	"github.com/qrug/go-exercises/api/add"

	"github.com/labstack/echo"
)

func main() {
	res := add.AddData()
	e := echo.New()
	fmt.Println(e)
	fmt.Println(res)
}
