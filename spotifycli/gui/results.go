package gui

import (
	"fmt"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type ResultsView struct {
	State   *core.State
	results *spotify.SearchResult
}

func (rv *ResultsView) render() error {
	maxX, maxY := rv.State.Gui.Size()
	if v, err := rv.State.Gui.SetView(resultsView, 15, 3, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " results (songs) "
		v.Wrap = true
	}
	return nil
}

func (rv *ResultsView) showResults(result *spotify.SearchResult) error {
	rv.results = result
	rv.State.Gui.Update(func(g *gocui.Gui) error {
		v, err := g.View(resultsView)
		if err != nil {
			return err
		}
		v.Clear()

		for _, track := range result.Tracks.Tracks {
			fmt.Fprintln(v, "♪ "+track.Artists[0].Name+" • "+track.Album.Name+" • "+track.Name)
		}
		return nil
	})
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (rv *ResultsView) playSong(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	_, xy := v.Origin()
	pos := cy + xy
	if pos < len(rv.results.Tracks.Tracks) {
		uris := make([]spotify.URI, 1)
		uris[0] = rv.results.Tracks.Tracks[pos].URI
		rv.State.Client.PlayOpt(&spotify.PlayOptions{
			URIs: uris,
		})
	}
	return nil
}

func (rv *ResultsView) bindKeys() error {
	if err := rv.State.Gui.SetKeybinding(resultsView, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := rv.State.Gui.SetKeybinding(resultsView, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := rv.State.Gui.SetKeybinding(resultsView, gocui.KeyEnter, gocui.ModNone, rv.playSong); err != nil {
		return err
	}
	return nil
}
