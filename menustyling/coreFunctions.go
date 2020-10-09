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
			if len(internalString) > menu.width - (menu.LineThickness * 4) - menu.WhiteSpace*2 {
				menu.width = len(internalString) + (menu.LineThickness * 4) + menu.WhiteSpace*2
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

	indexOfTextLine := 0
	indexOfTextChar := 0
	indexOfTextBlock := 0
	indexModForSeparator := 0
	indexOfSeparatorLine := 0

	for indexOfMenuLine := menu.LineThickness; indexOfMenuLine < menu.height - menu.LineThickness; indexOfMenuLine++ {
		indexOfTextChar = 0
		for indexOfMenuChar := menu.LineThickness*2 + menu.WhiteSpace; indexOfMenuChar < len(menu.MenuText[indexOfTextBlock][indexOfTextLine]) + menu.LineThickness*2 + menu.WhiteSpace; indexOfMenuChar++ {
			menu.menuArray[indexOfMenuLine][indexOfMenuChar] = string(menu.MenuText[indexOfTextBlock][indexOfTextLine][indexOfTextChar])
			indexOfTextChar++
		}


		if menu.WhiteSpace != 0 {
			for indexOfWhiteSpace := menu.LineThickness*2 + menu.WhiteSpace - 1; indexOfWhiteSpace > menu.LineThickness*2 - 1; indexOfWhiteSpace-- {
				menu.menuArray[indexOfMenuLine][indexOfWhiteSpace] = " "
			}
			for indexOfWhiteSpace := menu.LineThickness*2 + menu.WhiteSpace + len(menu.MenuText[indexOfTextBlock][indexOfTextLine]); indexOfWhiteSpace < menu.width - menu.LineThickness*2; indexOfWhiteSpace++ {
				menu.menuArray[indexOfMenuLine][indexOfWhiteSpace] = " "
			}
		}


		indexOfSeparatorLine++
		if menu.SeparatorLines && indexModForSeparator != len(menu.MenuText) - 1 && indexOfSeparatorLine == len(menu.MenuText[indexOfTextBlock]) {
			indexOfMenuLine++
			indexModForSeparator++
			indexOfSeparatorLine = 0
		}

		indexOfTextLine++
		if indexOfTextLine == len(menu.MenuText[indexOfTextBlock]) {
			indexOfTextLine = 0
		}

		indexOfTextBlock++
		if indexOfTextBlock == len(menu.MenuText) {
			indexOfTextBlock = 0
		}
	}
}