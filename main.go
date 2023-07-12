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
	window := tuikit.NewWindow()
	window.SetBorder(
		tuikit.BorderOpts{
			Style: tuikit.BorderStyle.Single,
			Color: "red",
		},
		tuikit.BorderOpts{
			Style: tuikit.BorderStyle.Single,
			Color: "red",
		},

		tuikit.BorderOpts{
			Style: tuikit.BorderStyle.Single,
			Color: "red",
		},

		tuikit.BorderOpts{
			Style: tuikit.BorderStyle.Single,
			Color: "red",
		},
	)
	window.SetSize(30, 10)
	window.SetMargin(2, 2, 2, 2)
	window.SetPadding(2, 2, 2, 2)
	// panic(
	// 	fmt.Sprintf(
	// 		"%v %v %v %v",
	// 		window.GetOuterWidth(),
	// 		window.GetOuterHeight(),
	// 		window.GetWidth(),
	// 		window.GetHeight(),
	// 	),
	// )
	screen.AddElement(window)
	// sel := tuikit.NewSelect([]string{"Choice 1", "Choice 2", "Choice 3"})
	// window.AddElement(sel)
	screen.Draw()

	<-(*screen.ExitChannel)
	screen.Close()
}
