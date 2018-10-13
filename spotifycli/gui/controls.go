package gui

import (
	"fmt"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
)

const (
	usernameRegex = "@([^\\s]+)"
)

/*
InputView shows message input
*/
type ControlsView struct {
	State *core.State
}

func (cv *ControlsView) render() error {
	_, maxY := cv.State.Gui.Size()
	if v, err := cv.State.Gui.SetView(controlsView, 0, maxY-3, 13, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " controls "
		v.Wrap = true
		fmt.Fprint(v, "  ← ▶ ⏹ →")

	}
	return nil
}

func (cv *ControlsView) bindKeys() error {
	return nil
}
