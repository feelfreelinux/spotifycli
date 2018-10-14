package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/gdamore/tcell"
	"github.com/jroimartin/gocui"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

const ()

/*
InputView shows message input
*/
type SearchView struct {
	State *core.State
	list  *tview.InputField
}

func (sv *SearchView) render() tview.Primitive {
	sv.list = tview.NewInputField()
	sv.list.SetBorder(true)
	sv.list.SetFieldBackgroundColor(tcell.ColorDefault)
	sv.list.SetTitle("search")
	sv.list.SetBackgroundColor(tcell.ColorDefault)
	return sv.list
}

func (sv *SearchView) search(g *gocui.Gui, v *gocui.View) error {
	v, err := g.View(searchView)
	if err != nil {
		return err
	}

	msg := v.Buffer()
	v.Clear()
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)

	go func() {
		result, err := sv.State.Client.Search(msg, spotify.SearchTypeTrack)
		if err != nil {
			return
		}

		sv.State.SearchResultsChan <- result
	}()

	return nil
}

func (sv *SearchView) bindKeys() error {
	if err := sv.State.Gui.SetKeybinding(searchView, gocui.KeyEnter, gocui.ModNone, sv.search); err != nil {
		return err
	}
	return nil
}
