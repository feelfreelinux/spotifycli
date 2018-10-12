package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
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
	if v, err := sv.State.Gui.SetView(searchView, 0, 3, maxX-1, 5); err != nil {
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

func (sv *SearchView) bindKeys() error {
	return nil
}
