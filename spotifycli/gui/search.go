package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

const ()

/*
InputView shows message input
*/
type SearchView struct {
	State *core.State
	list  *tview.InputField
}

func (sv *SearchView) render() tview.Primitive {
	sv.list = tview.NewInputField()
	sv.list.SetBorder(true)
	sv.list.SetFieldBackgroundColor(tcell.ColorDefault)
	sv.list.SetTitle("search")
	// sv.list.SetBackgroundColor(tcell.ColorDefault)
	sv.list.SetDoneFunc(func(key tcell.Key) {
		go func() {
			result, err := sv.State.Client.Search(sv.list.GetText(), spotify.SearchTypeTrack)
			if err != nil {
				return
			}

			sv.State.SearchResultsChan <- result
		}()
	})
	return sv.list
}

func (sv *SearchView) bindKeys() error {
	return nil
}
