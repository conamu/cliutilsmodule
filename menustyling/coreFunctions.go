package menustyling

func (menu *Menu) drawLine() {
	for _, v := range menu.MenuText {
		for _, w := range v {
			if len(w) >= menu.width {
				menu.width = len(w)
			}
		}
	}
	menu.line = make([]string, menu.width+ (menu.LineThickness* 2))
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
	for i := 0; i < menu.height; i++ {
		menu.menuArray[i] = menu.line
	}
}
