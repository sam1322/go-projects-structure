package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"snippetbox.sam1322/internal/models"
	"snippetbox.sam1322/internal/server"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

// var (
// 	database = os.Getenv("DB_DATABASE")
// 	password = os.Getenv("DB_PASSWORD")
// 	username = os.Getenv("DB_USERNAME")
// 	port     = os.Getenv("DB_PORT")
// 	host     = os.Getenv("DB_HOST")
// )

// // for a given DSN.
// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("pgx", dsn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("Connected to database")
// 	return db, nil
// }

func main() {

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// 	// Create a logger for writing error messages in the same way, but use stderr as
	// 	// the destination and use the log.Lshortfile flag to include the relevant
	// 	// file name and line number.
	// 	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	// 	db, err := openDB(connStr)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	defer db.Close()

	// 	app := &application{
	// 		errorLog: errorLog,
	// 		infoLog:  infoLog,
	// 		snippets: &models.SnippetModel{DB: db},
	// 	}

	// 	srv := &http.Server{
	// 		Addr:     ":8080",
	// 		ErrorLog: errorLog,
	// 		Handler:  app.routes(),
	// 	}

	// 	infoLog.Printf("Starting server on :8080")
	// 	// err := http.ListenAndServe(":8080", srv)
	// 	serverErr := srv.ListenAndServe()
	// 	errorLog.Fatal(serverErr)
	server := server.InitNewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	log.Println("Server started successfully!")
}
