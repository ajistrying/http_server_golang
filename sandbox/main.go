package main

import (
	"fmt"
	"log"

	"github.com/ajistrying/http_server_golang/internal/"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")
}
