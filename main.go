package main

import (
	"flag"

	"github.com/delley/gotiny/app"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port that server should listen on.")
	flag.Parse()
}
func main() {
	s := app.NewServer()

	s.Init(port)
	s.Start()
}
