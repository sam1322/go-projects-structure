package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"snippetbox.sam1322/internal/server"
)

func main() {

	server := server.InitNewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	log.Println("Server started successfully!")
}
