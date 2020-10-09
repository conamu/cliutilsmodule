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
			menu.menuArray[indexOfMenuLine][indexOfMenuChar] = menu.styleChar
		}
	}
}

func (menu *Menu) getWidth() {
	for _, internalStringArray := range menu.menuText {
		for _, internalString := range internalStringArray {
			if len(internalString) > menu.width - (menu.lineThickness* 4) - menu.whiteSpace*2 {
				menu.width = len(internalString) + (menu.lineThickness * 4) + menu.whiteSpace*2
			}
		}
	}
}

// getHeight calculates the height of the menu with the amount of lines in the menuText Array
func (menu *Menu) getHeight() {
	height := 0
	sepHeight := 0
	for _, stringArray := range menu.menuText {
		height += len(stringArray)
	}
	if menu.separatorLines {
		sepHeight = len(menu.menuText) - 1
	}
	menu.height = (height + (menu.lineThickness * 2)) + sepHeight
}

func (menu *Menu) putText() {

	indexOfTextLine := 0
	indexOfTextChar := 0
	indexOfTextBlock := 0
	indexModForSeparator := 0
	indexOfSeparatorLine := 0

	for indexOfMenuLine := menu.lineThickness; indexOfMenuLine < menu.height - menu.lineThickness; indexOfMenuLine++ {
		indexOfTextChar = 0
		for indexOfMenuChar := menu.lineThickness*2 + menu.whiteSpace; indexOfMenuChar < len(menu.menuText[indexOfTextBlock][indexOfTextLine]) + menu.lineThickness*2 + menu.whiteSpace; indexOfMenuChar++ {
			menu.menuArray[indexOfMenuLine][indexOfMenuChar] = string(menu.menuText[indexOfTextBlock][indexOfTextLine][indexOfTextChar])
			indexOfTextChar++
		}
		
		if menu.whiteSpace != 0 {
			for indexOfWhiteSpace := menu.lineThickness*2 + menu.whiteSpace - 1; indexOfWhiteSpace > menu.lineThickness*2 - 1; indexOfWhiteSpace-- {
				menu.menuArray[indexOfMenuLine][indexOfWhiteSpace] = " "
			}
			for indexOfWhiteSpace := menu.lineThickness*2 + menu.whiteSpace + len(menu.menuText[indexOfTextBlock][indexOfTextLine]); indexOfWhiteSpace < menu.width - menu.lineThickness*2; indexOfWhiteSpace++ {
				menu.menuArray[indexOfMenuLine][indexOfWhiteSpace] = " "
			}
		}

		indexOfSeparatorLine++
		indexOfTextLine++
		if menu.separatorLines && indexModForSeparator != len(menu.menuText) - 1 && indexOfSeparatorLine == len(menu.menuText[indexOfTextBlock]) {
			indexOfMenuLine++
			indexModForSeparator++
			indexOfSeparatorLine = 0
			indexOfTextBlock++
			indexOfTextLine = 0
		}
	}
}