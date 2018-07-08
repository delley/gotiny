package main

import "github.com/delley/gotiny/app"

func main() {
	s := app.NewServer()

	s.Init()
	s.Start()
}
