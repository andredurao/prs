package app

import (
	"github.com/jroimartin/gocui"
)

type App struct {
	Gui *gocui.Gui
}

func NewApp() (*App, error) {
	app := &App{}
	return app, nil
}
