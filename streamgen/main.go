package main

import (
	"github.com/clipperhouse/gen/typewriter"

	_ "github.com/Logiraptor/streams"
)

func main() {
	app, err := typewriter.NewApp("+stream")
	if err != nil {
		panic(err)
	}

	app.WriteAll()
}
