package gui

import (
	"fmt"
	"github.com/andredurao/prs/pkg/app"
	"github.com/jroimartin/gocui"
	"log"
)

func NewGui(app *app.App) {
	g := app.Gui
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", 'j', gocui.ModNone, cursorDown); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	width, height := g.Size()
	if v, err := g.SetView("header", 1, 1, width-1, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	if v, err := g.SetView("footer", 1, height-9, width-1, height-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "footer")
	}
	if v, err := g.SetView("container", 1, 4, width-1, height-10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack | gocui.AttrBold
		fmt.Fprintln(v, "container")
		fmt.Fprintln(v, "container")
		fmt.Fprintln(v, "container")
	}
	g.SetCurrentView("container")
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()
		max := 2
		if (oy + cy) < max {
			if err := v.SetCursor(cx, cy+1); err != nil {
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
