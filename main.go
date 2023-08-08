package main

import (
	"fmt"

	"github.com/hasban-fardani/user-CRUD-go/app"
)

func main() {
	fmt.Println("Hello World!")

	// test connect database
	app.Connect()
}
