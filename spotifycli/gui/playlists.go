package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type PlaylistsView struct {
	State   *core.State
	list    *tview.Table
	results *spotify.SimplePlaylistPage
}

func (pv *PlaylistsView) render() tview.Primitive {
	pv.list = tview.NewTable()
	pv.list.SetFixed(0, 0)
	pv.list.SetSeparator(tview.Borders.Vertical)
	pv.list.SetBorder(true)
	pv.list.SetSelectable(true, true)
	pv.list.SetTitle("playlist")
	// pv.list.SetBackgroundColor(tcell.ColorDefault)

	return pv.list
}

func (pv *PlaylistsView) showPlaylists(result *spotify.SimplePlaylistPage) error {
	go func() {
		pv.results = result
		pv.list.Clear()
		pv.list.SetSelectedFunc(pv.playPlaylist)
		for row, playlist := range result.Playlists {
			playlistCell := tview.NewTableCell(playlist.Name)

			playlistCell.SetExpansion(1)

			pv.list.SetCell(row+1, 0, playlistCell)
		}
		pv.list.ScrollToBeginning()

		pv.State.App.Draw()
	}()

	return nil
}

func (pv *PlaylistsView) playPlaylist(index int, column int) {

	uris := make([]spotify.URI, 1)
	uris[0] = pv.results.Playlists[index-1].URI
	pv.State.Client.PlayOpt(&spotify.PlayOptions{
		URIs: uris,
	})
}

func (cv *PlaylistsView) bindKeys() error {
	return nil
}
