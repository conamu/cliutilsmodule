package main

import (
	"fmt"
	"github.com/conamu/cliutilsmodule/menustyling"
	"github.com/conamu/cliutilsmodule/utils"
)

func main() {
	// test for the Utils getRandom Function
	fmt.Println(utils.GetRandom(20))

	// test for the menustyling
	var mmenu *menustyling.Menu = menustyling.CreateMenu()
	var smenu *menustyling.Menu = menustyling.CreateMenu()
	mmenu.StoreMenu("main")
	smenu.StoreMenu("secondary")

	// set the Text for the menu
	menuText := make([][]string, 3)
	menuLine := make([]string, 10)

	menuText[0] = menuLine
	menuText[1] = menuLine

	mmenu.MenuText = menuText
	mmenu.StyleChar = "@"
	mmenu.LineThickness = 2
	mmenu.SeparatorLines = true

	for i := 0; i < len(menuLine); i++ {
		menuLine[i] = "Q"
	}

	menustyling.DisplayStoredMenu("main")
	menustyling.DisplayStoredMenu("secondary")

}

