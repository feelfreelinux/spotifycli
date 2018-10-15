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
	State   *core.State
	results *spotify.SearchResult
	list    *tview.Table
}

func (rv *ResultsView) render() tview.Primitive {
	rv.list = tview.NewTable()
	rv.list.SetFixed(0, 0)
	rv.list.SetSeparator(tview.Borders.Vertical)
	rv.list.SetBorder(true)
	rv.list.SetSelectable(true, false)
	rv.list.SetTitle("results")
	// rv.list.SetBackgroundColor(tcell.ColorDarkSlateGray)
	rv.list.SetSelectedFunc(rv.playSong)
	return rv.list
}

func (rv *ResultsView) showResults(result *spotify.SearchResult) {
	rv.results = result
	rv.list.Clear()
	rv.list.SetCell(0, 2, tview.NewTableCell("[yellow]artist").SetSelectable(false))
	rv.list.SetCell(0, 0, tview.NewTableCell("[yellow]song").SetExpansion(3).SetSelectable(false))
	rv.list.SetCell(0, 1, tview.NewTableCell("[yellow]album").SetSelectable(false))
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

func (rv *ResultsView) playSong(index int, _ int) {
	uris := make([]spotify.URI, 1)
	uris[0] = rv.results.Tracks.Tracks[index-1].URI
	rv.State.Client.PlayOpt(&spotify.PlayOptions{
		URIs: uris,
	})

}

func (rv *ResultsView) bindKeys() error {
	return nil
}
