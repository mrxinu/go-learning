// SPDX-License-Identifier: Unlicense OR MIT

package main

// A simple Gio program. See https://gioui.org for more information.

import (
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	gofont.Register()
	th := material.NewTheme()
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e.Queue, e.Config, e.Size)
			l := material.H1(th, "Hello, Gio")
			maroon := color.RGBA{127, 0, 0, 255}
			l.Color = maroon
			l.Alignment = text.Middle
			l.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
