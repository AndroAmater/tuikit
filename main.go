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
	e.SetBorder(
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
	e.SetSize(30, 10)
	e.SetMargin(2, 2, 2, 2)
	e.SetPadding(2, 2, 2, 2)
	// panic(
	// 	fmt.Sprintf(
	// 		"%v %v %v %v",
	// 		e.GetOuterWidth(),
	// 		e.GetOuterHeight(),
	// 		e.GetWidth(),
	// 		e.GetHeight(),
	// 	),
	// )
	screen.AddElement(e)
	// sel := tuikit.NewSelect([]string{"Choice 1", "Choice 2", "Choice 3"})
	// e.AddElement(sel)
	screen.Draw()

	<-(*screen.ExitChannel)
	screen.Close()
}
