package main

import (
	"fmt"

	"router"
)

func main() {
	fmt.Println("Hello World")
	e := echo.New()

	e.Start(":8080")
}
