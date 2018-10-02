package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Recipes API")

	a := App{}
	a.Initialize(
		strings.TrimSpace(os.Getenv("APP_DB_HOST")),
		strings.TrimSpace(os.Getenv("APP_DB_PORT")),
		strings.TrimSpace(os.Getenv("APP_DB_USERNAME")),
		strings.TrimSpace(os.Getenv("APP_DB_PASSWORD")),
		strings.TrimSpace(os.Getenv("APP_DB_NAME")))

	a.Run(":8000")
}
