//go:build rp2040
// +build rp2040

package main

import (
	"image/color"
	"machine"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/adapter"
	"tinygo.org/x/drivers/ssd1289"
	"tinygo.org/x/drivers/touch"
	"tinygo.org/x/drivers/xpt2046"
)

func initHardware() (tinygui.Displayer, touch.Pointer) {

	/*
		bus := ssd1289.NewPinBus([16]machine.Pin{
			machine.GP4,
			machine.GP5,
			machine.GP6,
			machine.GP7,
			machine.GP8,
			machine.GP9,
			machine.GP10,
			machine.GP11,
			machine.GP12,
			machine.GP13,
			machine.GP14,
			machine.GP15,
			machine.GP16,
			machine.GP17,
			machine.GP18,
			machine.GP19,
		})
	*/
	bus := ssd1289.NewRP2040Bus(machine.GP4)

	utft := ssd1289.New(machine.GP0, machine.GP1, machine.GP2, machine.GP3, bus)
	utft.Configure()
	utft.FillDisplay(color.RGBA{250, 0, 0, 255})

	touch := xpt2046.New(machine.GP20, machine.GP21, machine.GP22, machine.GP26, machine.GP27)
	touch.Configure(&xpt2046.Config{
		Precision: 10,
	})
	rotated := adapter.UpsideDown(&utft, &touch)

	return &rotated, &rotated
}
