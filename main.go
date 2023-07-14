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
	e.SetPadding(1, 2, 1, 2)
	e.SetSize(50, 50)
	s := tuikit.NewSelect([]string{"Choice 1", "Choice 2", "Choice 3"}, []bool{})
	s.SetGrow(true, true)
	e.AddChild(s)
	screen.AddChild(e)
	screen.Draw()

	<-(*screen.ExitChannel)
	screen.Close()
}
