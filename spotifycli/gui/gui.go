package gui

import (
	"time"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

const (
	searchView    = "search"
	playbackView  = "playback"
	playlistsView = "playlists"
	controlsView  = "controls"
	resultsView   = "results"
)

/*
MainView holds reference for all views and renders them
*/
type MainView struct {
	search    *SearchView
	playback  *PlaybackView
	playlists *PlaylistsView
	controls  *ControlsView
	results   *ResultsView
	State     *core.State
}

func (mv *MainView) drawLayout() *tview.Flex {
	flex := tview.NewFlex()
	flex.AddItem(mv.playlists.render(), 0, 1, false)

	flex.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(mv.search.render(), 3, 1, true).
			AddItem(mv.results.render(), 0, 6, false), 0, 4, true)

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 8, true).
		AddItem(mv.playback.render(), 3, 0, false)
}

/*
CreateMainView creates MainView and all of its child views
*/
func CreateMainView(ui *tview.Application, client *spotify.Client) error {
	var state = &core.State{
		App:               ui,
		Client:            client,
		SearchResultsChan: make(chan *spotify.SearchResult),
	}
	var mainView = &MainView{
		State: state,

		results: &ResultsView{
			State: state,
		},
		search: &SearchView{
			State: state,
		},
		controls: &ControlsView{
			State: state,
		},
		playlists: &PlaylistsView{
			State: state,
		},
		playback: &PlaybackView{
			State: state,
		},
	}

	mainView.setHandlers()

	ui.SetRoot(mainView.drawLayout(), true)
	return nil
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

	/*go func() {
		for {
			search := <-mv.State.SearchResultsChan
			mv.results.showResults(search)
		}
	}()*/

	go func() {
		results, err := mv.State.Client.CurrentUsersPlaylists()
		if err != nil {
			return
		}

		mv.playlists.showPlaylists(results)
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

	if err := mv.results.bindKeys(); err != nil {
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
	case searchView:
		g.SetCurrentView(resultsView)

	case resultsView:
		g.SetCurrentView(playlistsView)

	case playlistsView:
		g.SetCurrentView(searchView)
	}
	return nil
}
