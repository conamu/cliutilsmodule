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

// Function to put Whitespaces, Separator lines and the Menu text into the Menu field
func (menu *Menu) putText() {

	// Set some local variables to keep track of important positions and configurations.
	// Keep track of the current Textblock index
	indexOfTextBlock := 0
	// Keep track of the current Text Line of the current text block.
	indexOfTextLine := 0
	// Keep track of the current Character of the current text line.
	indexOfTextChar := 0
	// Mod value to insert the Separator correctly
	indexModForSeparator := 0
	// The index of text line between to text blocks where the line is placed.
	indexOfSeparatorLine := 0

	// Set first line at the line thickness value, as we dont need to change anything before that.
	for indexOfMenuLine := menu.lineThickness; indexOfMenuLine < menu.height - menu.lineThickness; indexOfMenuLine++ {
		// At the beginning off each new text line, set the character index back to 0.
		indexOfTextChar = 0
		// Insert the Text inside the MenuText array into the Menu field.
		for indexOfMenuChar := menu.lineThickness*2 + menu.whiteSpace; indexOfMenuChar < len(menu.menuText[indexOfTextBlock][indexOfTextLine]) + menu.lineThickness*2 + menu.whiteSpace; indexOfMenuChar++ {
			menu.menuArray[indexOfMenuLine][indexOfMenuChar] = string(menu.menuText[indexOfTextBlock][indexOfTextLine][indexOfTextChar])
			indexOfTextChar++
		}

		// If some number of Whitespace is specified, replace the space before and after the text with spaces.
		if menu.whiteSpace != 0 {
			// This replaces characters before the text.
			for indexOfWhiteSpace := menu.lineThickness*2 + menu.whiteSpace - 1; indexOfWhiteSpace > menu.lineThickness*2 - 1; indexOfWhiteSpace-- {
				menu.menuArray[indexOfMenuLine][indexOfWhiteSpace] = " "
			}
			// this replaces characters after the text.
			for indexOfWhiteSpace := menu.lineThickness*2 + menu.whiteSpace + len(menu.menuText[indexOfTextBlock][indexOfTextLine]); indexOfWhiteSpace < menu.width - menu.lineThickness*2; indexOfWhiteSpace++ {
				menu.menuArray[indexOfMenuLine][indexOfWhiteSpace] = " "
			}
		}

		// At the end of the loop, add 1 to the text line index and separator line index.
		indexOfSeparatorLine++
		indexOfTextLine++
		// if the separator line option is true, hop over one text line to leave a blank separator line.
		if menu.separatorLines && indexModForSeparator != len(menu.menuText) - 1 && indexOfSeparatorLine == len(menu.menuText[indexOfTextBlock]) {
			indexOfMenuLine++
			indexModForSeparator++
			indexOfSeparatorLine = 0
			indexOfTextBlock++
			indexOfTextLine = 0
		} else if !menu.separatorLines && indexModForSeparator != len(menu.menuText) - 1 && indexOfSeparatorLine == len(menu.menuText[indexOfTextBlock]) {
			indexOfSeparatorLine = 0
			indexOfTextBlock++
			indexOfTextLine = 0
		}
	}
}