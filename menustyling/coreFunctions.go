package menustyling

func (menu *Menu) drawLine() {
	for _, v := range menu.MenuText {
		for _, w := range v {
			if len(w) >= menu.Width {
				menu.Width = len(w)
			}
		}
	}
	menu.Line = make([]string, (menu.Width + (menu.LineThicness * 2)))
	for i := 0; i < menu.Width; i++ {
		menu.Line[i] = menu.StyleChar
	}
}

func (menu *Menu) getHeight() {
	var length int = 0
	sepLen := 0
	for _, v := range menu.MenuText {
		length += len(v)
	}

	if menu.SeparatorLines == true {
		sepLen = len(menu.MenuText) - 1
	}

	menu.Height = (length + (menu.LineThicness * 2)) + sepLen
}

func (menu *Menu) drawField() {
	menu.getHeight()
	menu.drawLine()
	for i := 0; i < menu.Height; i++ {
		menu.MenuArray[i] = menu.Line
	}
}
