package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

func setKeybindings(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", 'j', gocui.ModNone, cursorDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", gocui.KeyArrowDown, gocui.ModNone,
		cursorDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", 'k', gocui.ModNone, cursorUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", gocui.KeyArrowUp, gocui.ModNone,
		cursorUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("container", '/', gocui.ModNone, showSearch); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, search); err != nil {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		rows := len(*renderResult.RenderMap())
		if (oy + cy) < (rows - 1) {
			if err := v.SetCursor(cx, cy+1); err != nil {
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func showSearch(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("search", maxX/2-30, maxY/2-1, maxX/2+30, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		if _, err := g.SetCurrentView("search"); err != nil {
			return err
		}
	}
	return nil
}

func search(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("search"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("default"); err != nil {
		return err
	}
	return nil
}
