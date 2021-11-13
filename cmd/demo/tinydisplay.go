//go:build !tinygo
// +build !tinygo

package main

import (
	"log"

	"github.com/sago35/tinydisplay"
	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/adapter"
	"tinygo.org/x/drivers/touch"
)

var tempDelta = float32(0.1)
var lastTemp = float32(24)

func ReadTemp() float32 {
	lastTemp = lastTemp + tempDelta
	if lastTemp > 25 || lastTemp < 18.5 {
		tempDelta = -tempDelta
	}

	return lastTemp
}

func initHardware() (tinygui.Displayer, touch.Pointer) {
	display, err := tinydisplay.NewClient("", 9812, 240, 320)
	if err != nil {
		log.Fatal(err)
	}

	fillRectAdapter := adapter.NewTinyDisplayAdapter(display)

	return &fillRectAdapter, display
}
