package menustyling

import (
	"fmt"
	"strings"
)

// MenuDataStore is a Map which stores multiple menus to use later
var MenuDataStore MenuStore

// CreateMenu Creates and Initialises a new Menu fo you to Customize.
func CreateMenu() *Menu {
	var menu Menu
	var pmenu = &menu
	pmenu.Height = 0
	pmenu.LineThicness = 0
	pmenu.Width = 0
	pmenu.SeparatorLines = false
	pmenu.StyleChar = ""
	pmenu.Line = nil
	pmenu.MenuText = nil
	pmenu.MenuArray = nil
	pmenu.Menu = ""

	return pmenu
}

// StoreMenu stores you menu object into a custom map, indexed by a string identifier
func (menu *Menu) StoreMenu(identifier string) {
	MenuDataStore[identifier] = menu
}

// DisplayStoredMenu displays a Menu form the MenuStore base on the identifier
func DisplayStoredMenu(indentifier string) *Menu {
	return MenuDataStore[indentifier]
}

// DisplayMenu displays the menu.
func (menu *Menu) DisplayMenu() {
	menu.drawField()
	stringArray := make([]string, menu.Height)

	for i := 0; i < len(menu.MenuArray); i++ {
		stringArray[i] = strings.Join(menu.MenuArray[i], "")
	}

	menu.Menu = strings.Join(stringArray, "\n")

	fmt.Println(menu.Menu)
}
