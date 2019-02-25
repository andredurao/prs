package app

import (
	"github.com/jroimartin/gocui"
	"log"
)

type App struct {
	Gui      *gocui.Gui
	Messages chan string
}

func NewApp() (*App, error) {
	app := &App{}
	gui, err := gocui.NewGui(gocui.OutputNormal)
	gui.Cursor = true
	if err != nil {
		log.Panicln(err)
	}
	app.Gui = gui
	app.Messages = make(chan string)
	return app, nil
}
