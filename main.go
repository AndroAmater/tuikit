package main

import (
	"tuikit/tuikit"
)

func main() {
	screen, err := tuikit.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Screen.Init(); err != nil {
		panic(err)
	}

	screen.Screen.Clear()
	e := tuikit.NewElement()
	e.SetFullBorder(
		tuikit.BorderStyle.Single,
		"white",
	)
	e.SetMargin(2, 2, 2, 2)
	e.SetPadding(2, 2, 2, 2)
	e.SetGrow(true, true)
	e2 := tuikit.NewElement()
	e2.SetFullBorder(
		tuikit.BorderStyle.Single,
		"white",
	)
	e2.SetPadding(2, 2, 2, 2)
	e2.SetGrow(true, true)
	e3 := tuikit.NewElement()
	e3.SetFullBorder(
		tuikit.BorderStyle.Single,
		"white",
	)
	e3.SetMargin(2, 2, 2, 2)
	e3.SetPadding(2, 2, 2, 2)
	e3.SetGrow(true, false)
	e3.SetHeight(40)
	e4 := tuikit.NewElement()
	e4.SetFullBorder(
		tuikit.BorderStyle.Single,
		"white",
	)
	e4.SetGrow(true, true)
	e4.SetContentDirection(tuikit.ContainerDirection.Column)
	e5 := tuikit.NewElement()
	e5.SetFullBorder(
		tuikit.BorderStyle.Single,
		"white",
	)
	e5.SetGrow(true, true)
	e4.AddChild(e3)
	e4.AddChild(e5)
	e.AddChild(e2)
	e.AddChild(e4)
	screen.AddChild(e)
	// sel := tuikit.NewSelect([]string{"Choice 1", "Choice 2", "Choice 3"})
	// e.AddChild(sel)
	screen.Draw()

	<-(*screen.ExitChannel)
	screen.Close()
}
