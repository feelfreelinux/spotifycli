package core

import (
	"github.com/jroimartin/gocui"
	"github.com/zmb3/spotify"
)

/*
State holds reference of current application state (selected channel, etc)
*/
type State struct {
	Gui    *gocui.Gui
	Client *spotify.Client
}
