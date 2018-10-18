package gui

import (
	"github.com/feelfreelinux/spotifycli/spotifycli/core"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

/*
InputView shows message input
*/
type DevicesView struct {
	State *core.State
	list  *tview.Table
}

func (dv *DevicesView) render() tview.Primitive {
	dv.list = tview.NewTable()
	dv.list.SetTitle("devices")
	dv.list.SetSelectable(true, false)

	dv.list.SetBorder(true)

	go func() {
		playerDevices, err := dv.State.Client.PlayerDevices()
		if err != nil {
			return
		}
		dv.showDevices(playerDevices)
	}()
	return dv.list
}

func (dv *DevicesView) showDevices(devices []spotify.PlayerDevice) error {
	dv.list.Clear()
	dv.list.SetSelectedFunc(
		func(index int, _ int) {
			dv.State.Client.TransferPlayback(devices[index].ID, true)
		},
	)
	for index, device := range devices {
		dv.list.SetCellSimple(index, 0, device.Name)
	}
	dv.State.App.Draw()
	return nil
}

func (dv *DevicesView) bindKeys() error {
	return nil
}
