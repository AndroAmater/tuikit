package main

import (
	"mtd-staging-migrator/tuikit"
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
	window := tuikit.NewWindow()
	screen.AddElement(window)
	sel := tuikit.NewSelect([]string{"Choice 1", "Choice 2", "Choice 3"})
	window.AddElement(sel)
	screen.Draw()

	<-(*screen.ExitChannel)
	screen.Close()
}
