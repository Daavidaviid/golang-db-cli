package main

import (
	"blogCLI/app"
	"blogCLI/database"
	"fmt"
)

func main() {
	err := app.LaunchCLI()
	if err != nil {
		fmt.Println(err)
	}

	// To close database connection
	if database.DB != nil {
		database.DB.Close()
	}
}
