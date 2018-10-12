package gui

import (
	"time"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
	"github.com/zmb3/spotify"
)

const (
	searchView   = "search"
	playbackView = "playback"
)

/*
MainView holds reference for all views and renders them
*/
type MainView struct {
	search   *SearchView
	playback *PlaybackView
	State    *core.State
}

func (mv *MainView) layout(g *gocui.Gui) error {
	if err := mv.search.render(); err != nil {
		return err
	}

	if err := mv.playback.render(); err != nil {
		return err
	}

	return nil
}

/*
CreateMainView creates MainView and all of its child views
*/
func CreateMainView(ui *gocui.Gui, client *spotify.Client) error {
	ui.Cursor = true
	var state = &core.State{
		Gui:    ui,
		Client: client,
	}
	var mainView = &MainView{
		State: state,
		search: &SearchView{
			State: state,
		},
		playback: &PlaybackView{
			State: state,
		},
	}
	ui.SetManagerFunc(mainView.layout)
	err := mainView.bindKeys()

	mainView.setHandlers()
	return err
}

func (mv *MainView) setHandlers() error {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				go func() {
					currentlyPlaying, err := mv.State.Client.PlayerCurrentlyPlaying()
					if err != nil {
						return
					}
					mv.playback.drawPlaybackState(currentlyPlaying)
				}()
			}
		}
	}()
	return nil
}

func (mv *MainView) bindKeys() error {
	if err := mv.State.Gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, mv.quit); err != nil {
		return err
	}

	if err := mv.State.Gui.SetKeybinding("", gocui.KeyCtrlSpace, gocui.ModNone, changeScreenFocus); err != nil {
		return err
	}

	if err := mv.search.bindKeys(); err != nil {
		return err
	}

	if err := mv.search.bindKeys(); err != nil {
		return err
	}
	return nil
}

func (mv *MainView) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func changeScreenFocus(g *gocui.Gui, v *gocui.View) error {
	switch g.CurrentView().Name() {

	}
	return nil
}
