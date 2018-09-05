package app

import (
	"github.com/jroimartin/gocui"
	"log"
)

type App struct {
	Gui *gocui.Gui
}

func NewApp() (*App, error) {
	app := &App{}
	gui, err := gocui.NewGui(gocui.OutputNormal)
	gui.Cursor = true
	if err != nil {
		log.Panicln(err)
	}
	app.Gui = gui
	return app, nil
}
