package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type ResultsView struct {
	State         *core.State
	results       *spotify.SearchResult
	albumResults  *spotify.SimpleTrackPage
	artistResults *spotify.SimpleAlbumPage
	currentAlbum  *spotify.SimpleAlbum
	currentArtist *spotify.SimpleArtist
	list          *tview.Table
}

func (rv *ResultsView) render() tview.Primitive {
	rv.list = tview.NewTable()
	rv.list.SetFixed(0, 0)
	rv.list.SetSeparator(tview.Borders.Vertical)
	rv.list.SetBorder(true)
	rv.list.SetSelectable(true, true)
	rv.list.SetTitle("results")
	// rv.list.SetBackgroundColor(tcell.ColorDarkSlateGray)
	return rv.list
}

func (rv *ResultsView) showResults(result *spotify.SearchResult) {
	rv.results = result
	rv.list.Clear()
	rv.list.SetSelectedFunc(rv.playSong)
	rv.list.SetTitle("results")
	rv.list.SetCell(0, 2, tview.NewTableCell("[yellow]artist").SetSelectable(false).SetAlign(tview.AlignCenter))
	rv.list.SetCell(0, 0, tview.NewTableCell("[yellow]song").SetExpansion(3).SetSelectable(false).SetAlign(tview.AlignCenter))
	rv.list.SetCell(0, 1, tview.NewTableCell("[yellow]album").SetSelectable(false).SetAlign(tview.AlignCenter))
	for row, track := range result.Tracks.Tracks {
		artistCell := tview.NewTableCell(track.Artists[0].Name)
		songCell := tview.NewTableCell(track.Name)
		albumCell := tview.NewTableCell(track.Album.Name)

		songCell.SetExpansion(3)

		rv.list.SetCell(row+1, 2, artistCell)
		rv.list.SetCell(row+1, 0, songCell)
		rv.list.SetCell(row+1, 1, albumCell)
	}
	rv.list.ScrollToBeginning()

	rv.State.App.Draw()
}

func (rv *ResultsView) showAlbum(album *spotify.SimpleAlbum) {
	go func() {
		result, err := rv.State.Client.GetAlbumTracks(album.ID)
		if err != nil {
			return
		}
		rv.currentAlbum = album
		rv.albumResults = result
		rv.list.SetSelectedFunc(rv.playAlbum)
		rv.list.SetTitle(album.Name)
		rv.list.Clear()
		rv.list.SetCell(0, 2, tview.NewTableCell("[yellow]artist").SetSelectable(false).SetAlign(tview.AlignCenter))
		rv.list.SetCell(0, 0, tview.NewTableCell("[yellow]song").SetExpansion(3).SetSelectable(false).SetAlign(tview.AlignCenter))
		rv.list.SetCell(0, 1, tview.NewTableCell("[yellow]album").SetSelectable(false).SetAlign(tview.AlignCenter))
		for row, track := range result.Tracks {
			artistCell := tview.NewTableCell(track.Artists[0].Name)
			songCell := tview.NewTableCell(track.Name)
			albumCell := tview.NewTableCell(album.Name)

			songCell.SetExpansion(3)

			rv.list.SetCell(row+1, 2, artistCell)
			rv.list.SetCell(row+1, 0, songCell)
			rv.list.SetCell(row+1, 1, albumCell)
		}
		rv.list.ScrollToBeginning()

		rv.State.App.Draw()
	}()

}

func (rv *ResultsView) showArtist(artist *spotify.SimpleArtist) {
	go func() {
		result, err := rv.State.Client.GetArtistAlbums(artist.ID)
		if err != nil {
			return
		}
		rv.currentArtist = artist
		rv.artistResults = result
		rv.list.SetSelectedFunc(func(index int, _ int) {
			rv.showAlbum(&result.Albums[index-1])
		})
		rv.list.SetTitle(artist.Name)
		rv.list.Clear()
		rv.list.SetCell(0, 1, tview.NewTableCell("[yellow]artist").SetSelectable(false).SetAlign(tview.AlignCenter))
		rv.list.SetCell(0, 0, tview.NewTableCell("[yellow]album").SetSelectable(false).SetExpansion(1).SetAlign(tview.AlignCenter))
		for row, album := range result.Albums {
			artistCell := tview.NewTableCell(artist.Name)
			albumCell := tview.NewTableCell(album.Name)

			albumCell.SetExpansion(1)

			rv.list.SetCell(row+1, 1, artistCell)
			rv.list.SetCell(row+1, 0, albumCell)
		}
		rv.list.ScrollToBeginning()

		rv.State.App.Draw()
	}()

}

func (rv *ResultsView) playSong(index int, column int) {
	if column == 1 {
		rv.showAlbum(&rv.results.Tracks.Tracks[index-1].Album)
		return
	}

	if column == 2 {
		rv.showArtist(&rv.results.Tracks.Tracks[index-1].Artists[0])
	}

	uris := make([]spotify.URI, 1)
	uris[0] = rv.results.Tracks.Tracks[index-1].URI
	rv.State.Client.PlayOpt(&spotify.PlayOptions{
		URIs: uris,
	})
}

func (rv *ResultsView) playAlbum(index int, column int) {
	if column == 2 {
		rv.showArtist(&rv.albumResults.Tracks[index-1].Artists[0])
	}

	if column == 1 {
		rv.showAlbum(rv.currentAlbum)
	}

	uris := make([]spotify.URI, 1)
	uris[0] = rv.albumResults.Tracks[index-1].URI
	rv.State.Client.PlayOpt(&spotify.PlayOptions{
		PlaybackOffset: &spotify.PlaybackOffset{
			URI: uris[0],
		},
		PlaybackContext: &rv.currentAlbum.URI,
	})
}

func (rv *ResultsView) bindKeys() error {
	return nil
}
