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
type PlaylistsView struct {
	State *core.State
}

func (pv *PlaylistsView) render() error {
	_, maxY := pv.State.Gui.Size()
	if v, err := pv.State.Gui.SetView(playlistsView, 0, 0, 13, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " playlists "
	}
	return nil
}

func (pv *PlaylistsView) showPlaylists(playlists *spotify.SimplePlaylistPage) error {
	pv.State.Gui.Update(func(g *gocui.Gui) error {
		v, err := g.View(playlistsView)
		if err != nil {
			return err
		}
		v.Clear()

		for _, playlist := range playlists.Playlists {
			fmt.Fprintln(v, playlist.Name)
		}

		return nil
	})
	return nil
}

func (cv *PlaylistsView) bindKeys() error {
	return nil
}
