package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/jroimartin/gocui"
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
		v.Wrap = true
	}
	return nil
}

/* func (pv *PlaylistsView) showPlaylists(state *spotify.GetPlaylists) error {
	pv.State.Gui.Update(func(g *gocui.Gui) error {
		fmt.Printf("ddd")
		maxX, _ := pv.State.Gui.Size()
		v, err := g.View(PlaylistsView)
		if err != nil {
			return err
		}
		v.Clear()

		var spotfiy map[string]interface{}

		func (spotify[playlists].(map[string]interface{})) {
			fmt.Printf(playlists[name])
		}

		return nil
	})
} */

func (cv *PlaylistsView) bindKeys() error {
	return nil
}
