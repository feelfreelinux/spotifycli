package gui

import (
	"fmt"
	"strings"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type PlaybackView struct {
	State *core.State
}

func (pv *PlaybackView) render() error {
	maxX, maxY := pv.State.Gui.Size()
	if v, err := pv.State.Gui.SetView(playbackView, 15, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " playback "
		v.Wrap = true
		if _, err := pv.State.Gui.SetCurrentView(playbackView); err != nil {
			return err
		}
	}
	return nil
}

func (cv *PlaybackView) drawPlaybackState(state *spotify.CurrentlyPlaying) error {
	cv.State.Gui.Update(func(g *gocui.Gui) error {
		maxX, _ := cv.State.Gui.Size()

		v, err := g.View(playbackView)
		if err != nil {
			return err
		}
		v.Clear()

		v.Title = " " + state.Item.Artists[0].Name + ": " + state.Item.Name + " "
		rep := int(float64(float64(state.Progress)/float64(state.Item.Duration)) * float64(maxX-18))
		fmt.Fprint(v, strings.Repeat("▒", rep))
		return nil
	})
	return nil
}

func (pv *PlaybackView) bindKeys() error {
	return nil
}
