package gui

import (
	"fmt"
	"github.com/andredurao/prs/pkg/app"
	"github.com/andredurao/prs/pkg/renderer"
	"github.com/jroimartin/gocui"
	"log"
	"time"
)

var renderResult renderer.Renderer

func NewGui(app *app.App) {
	g := app.Gui
	defer g.Close()

	g.SetManagerFunc(layout)

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

	messages := make(chan string)
	go func() {
		getPullRequests(messages, g)
		for {
			updateView(g, "container", <-messages)
		}
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func getPullRequests(c chan string, g *gocui.Gui) {
	time.AfterFunc(50*time.Millisecond, func() {
		clearView(g, "container")
		FilterRows()
		for _, row := range *renderResult.RenderMap() {
			c <- fmt.Sprintf(row.Row)
		}
	})
}

func clearView(g *gocui.Gui, view string) {
	v, err := g.View(view)
	if err != nil {
		log.Println("Cannot get output view:", err)
	}
	v.Clear()
}

func updateView(g *gocui.Gui, view string, content string) {
	v, err := g.View(view)
	if err != nil {
		log.Println("Cannot get output view:", err)
	}
	fmt.Fprintln(v, content)
	g.Update(func(g *gocui.Gui) error {
		return nil
	})
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
		help := `
q: quit ?: help spacebar: open PR URL
return: git checkout "/": search
j/↓: cursor down k/↑: cursor up
`
		fmt.Fprintln(v, help)
	}
	if v, err := g.SetView("container", 1, 4, width-1, height-10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack | gocui.AttrBold
		fmt.Fprintln(v, "Loading...")
	}
	g.SetCurrentView("container")
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
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
