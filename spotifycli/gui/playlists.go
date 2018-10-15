package gui

import (
	"fmt"

	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type PlaylistsView struct {
	State *core.State
	list  *tview.TextView
}

func (pv *PlaylistsView) render() tview.Primitive {
	pv.list = tview.NewTextView()
	pv.list.SetTitle("playlists")
	pv.list.SetTitleAlign(tview.AlignCenter)
	pv.list.SetBorder(true)
	pv.list.SetScrollable(true)
	// pv.list.SetBackgroundColor(tcell.ColorDefault)
	return pv.list
}

func (pv *PlaylistsView) showPlaylists(playlists *spotify.SimplePlaylistPage) error {
	for _, playlist := range playlists.Playlists {
		fmt.Fprintln(pv.list, playlist.Name)
	}

	pv.State.App.Draw()
	return nil
}

func (cv *PlaylistsView) bindKeys() error {
	return nil
}
