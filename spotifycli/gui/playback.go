package gui

import (
	"fmt"
	"strings"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type PlaybackView struct {
	State *core.State
	bar   *tview.TextView
}

func (pv *PlaybackView) render() tview.Primitive {
	pv.bar = tview.NewTextView()
	pv.bar.SetTitle("playback")
	pv.bar.SetBackgroundColor(tcell.ColorDefault)
	pv.bar.SetBorder(true)
	return pv.bar
}

func (cv *PlaybackView) drawPlaybackState(state *spotify.CurrentlyPlaying) error {
	_, _, width, _ := cv.bar.GetInnerRect()
	cv.bar.Clear()
	cv.bar.SetTitle(" [red]" + state.Item.Artists[0].Name + "[grey] - [blue]" + state.Item.Name + " ")
	rep := int(float64(float64(state.Progress)/float64(state.Item.Duration)) * float64(width))

	fmt.Fprint(cv.bar, strings.Repeat("▒", rep))
	cv.State.App.Draw()
	return nil
}

func (pv *PlaybackView) bindKeys() error {
	return nil
}
