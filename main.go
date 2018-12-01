package main

import (
	"github.com/andredurao/prs/pkg/app"
	"github.com/andredurao/prs/pkg/gui"
)

func main() {
	app, _ := app.NewApp()
	gui.NewGui(app)
}
