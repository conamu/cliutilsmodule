package menustyling

import "fmt"

//TODO: Fix Line Rendering, string length is not beeing calculated correctly.
func (menu *Menu) drawLine() {
	for _, internalStringArray := range menu.MenuText {
		for _, internalString := range internalStringArray {
			fmt.Println("DEBUG:", len(internalString))
			fmt.Println("DEBUG:", len(internalStringArray))
			if len(internalString) >= menu.width {
				fmt.Println("DEBUG: Successfully updated Width:", menu.width)
				menu.width = len(internalString)
			}
		}
	}

	menu.line = make([]string, menu.width + (menu.LineThickness* 2))
	for i := 0; i < menu.width; i++ {
		menu.line[i] = menu.StyleChar
	}
}

func (menu *Menu) getHeight() {
	length := 0
	sepLen := 0
	for _, v := range menu.MenuText {
		length += len(v)
	}

	if menu.SeparatorLines == true {
		sepLen = len(menu.MenuText) - 1
	}

	menu.height = (length + (menu.LineThickness * 2)) + sepLen
}

func (menu *Menu) drawField() {
	menu.getHeight()
	menu.drawLine()
	menu.menuArray = make([][]string, menu.height)
	for i := 0; i < menu.height; i++ {
		menu.menuArray[i] = menu.line
	}
}
