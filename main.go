package main

import (
	"mtd-staging-migrator/tuikit"
)

var (
	choices  = []string{"Choice 1", "Choice 2", "Choice 3"}
	selected = make([]bool, len(choices))
	cursor   int
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
	sel := tuikit.NewSelect(choices, selected)
	screen.AddElement(sel)
	// window := tuikit.NewWindow(0, 0, 20, 20, false, false)
	// screen.AddElement(window)
	screen.Draw()

	<-(*screen.ExitChannel)
	screen.Close()
}
