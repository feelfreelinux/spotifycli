package gui

import (
	"time"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/gdamore/tcell"
	"github.com/jroimartin/gocui"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

/*
MainView holds reference for all views and renders them
*/
type MainView struct {
	search    *SearchView
	playback  *PlaybackView
	playlists *PlaylistsView
	devices   *DevicesView
	results   *ResultsView
	State     *core.State
}

func (mv *MainView) drawLayout() *tview.Flex {
	flex := tview.NewFlex()
	flex.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(mv.playlists.render(), 0, 5, false).
			AddItem(mv.devices.render(), 0, 2, false), 0, 1, false)

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
		playlists: &PlaylistsView{
			State: state,
		},
		playback: &PlaybackView{
			State: state,
		},
		devices: &DevicesView{
			State: state,
		},
	}

	mainView.setHandlers()
	mainView.bindKeys()

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

	go func() {
		for {
			search := <-mv.State.SearchResultsChan
			mv.results.showResults(search)
		}
	}()

	go func() {
		results, err := mv.State.Client.CurrentUsersPlaylists()
		if err != nil {
			return
		}

		mv.playlists.showPlaylists(results)
	}()
	return nil
}

func (mv *MainView) bindKeys() {

	mv.State.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if mv.State.App.GetFocus() == mv.search.list {
				mv.State.App.SetFocus(mv.results.list)
				mv.State.App.Draw()
			} else if mv.State.App.GetFocus() == mv.results.list {
				mv.State.App.SetFocus(mv.devices.list)
				mv.State.App.Draw()
			} else if mv.State.App.GetFocus() == mv.devices.list {
				mv.State.App.SetFocus(mv.search.list)
				mv.State.App.Draw()
			}
		}
		return event
	})
}

func (mv *MainView) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
