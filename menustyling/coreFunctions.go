package menustyling

// drawField generates decoration lines for the menu.
// It will generate a line based on the longest string and lineThickness
// Then it will fill out a 2 dimensional field with those characters.
func (menu *Menu) drawField() {
	menu.getWidth()
	menu.getHeight()
	menu.menuArray = make([][]string, menu.height)

	for indexOfMenuLine := 0; indexOfMenuLine < menu.height; indexOfMenuLine++ {
		menu.menuArray[indexOfMenuLine] = make([]string, menu.width)
		for indexOfMenuChar := 0; indexOfMenuChar < menu.width; indexOfMenuChar++ {
			menu.menuArray[indexOfMenuLine][indexOfMenuChar] = menu.StyleChar
		}
	}
}

func (menu *Menu) getWidth() {
	for _, internalStringArray := range menu.MenuText {
		for _, internalString := range internalStringArray {
			if len(internalString) >= menu.width {
				menu.width = len(internalString) + (menu.LineThickness * 2)
			}
		}
	}
}

// getHeight calculates the height of the menu with the amount of lines in the menuText Array
func (menu *Menu) getHeight() {
	height := 0
	sepHeight := 0
	for _, stringArray := range menu.MenuText {
		height += len(stringArray)
	}
	if menu.SeparatorLines {
		sepHeight = len(menu.MenuText) - 1
	}
	menu.height = (height + (menu.LineThickness * 2)) + sepHeight
}

func (menu *Menu) putText() {
	
}
