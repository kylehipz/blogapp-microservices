package internal

import "os"

var (
	API_SERVER_PORT = os.Getenv("PORT")
	DATABASE_URL    = os.Getenv("DATABASE_URL")
)
