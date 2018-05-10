package main

import (
	"fmt"

	"github.com/urfave/negroni"
)

func main() {
	fmt.Println("vim-go")
	app := negroni.Classic()
	app.Run(":8080")
}
