package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
)

/*
InputView shows message input
*/
type ResultsView struct {
	State *core.State
}

func (rv *ResultsView) render() error {
	maxX, maxY := rv.State.Gui.Size()
	if v, err := rv.State.Gui.SetView(resultsView, 15, 3, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " results "
		v.Wrap = true
	}
	return nil
}

func (rv *ResultsView) bindKeys() error {
	return nil
}
