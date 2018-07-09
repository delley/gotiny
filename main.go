package main

import (
	"flag"
	"os"

	"mod/github.com/joho/godotenv@v1.2.0"

	"github.com/delley/gotiny/app"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port that server should listen on.")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}
func main() {
	s := app.NewServer()

	s.Init(port)
	s.Start()
}
