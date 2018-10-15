package core

import (
	"github.com/rivo/tview"
)

type View interface {
	getView() tview.Primitive
	createHandlers() error
}
