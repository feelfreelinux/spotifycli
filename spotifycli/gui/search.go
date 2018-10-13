package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
	"github.com/zmb3/spotify"
)

const ()

/*
InputView shows message input
*/
type SearchView struct {
	State *core.State
}

func (sv *SearchView) render() error {
	maxX, _ := sv.State.Gui.Size()
	if v, err := sv.State.Gui.SetView(searchView, 15, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = false
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " search "
		v.Wrap = true
		if _, err := sv.State.Gui.SetCurrentView(searchView); err != nil {
			return err
		}
	}
	return nil
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
