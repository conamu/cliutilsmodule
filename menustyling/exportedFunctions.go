package menustyling

import (
	"fmt"
	"strings"
)

// menuDataStore is a Map which stores multiple menus to use later
var menuDataStore menuStore

// CreateMenu Creates and Initialises a new Menu fo you to Customize.
func CreateMenu() *Menu {
	var menu Menu
	var pMenu = &menu
	pMenu.height = 0
	pMenu.LineThickness = 0
	pMenu.width = 0
	pMenu.SeparatorLines = false
	pMenu.StyleChar = ""
	pMenu.line = nil
	pMenu.MenuText = nil
	pMenu.menuArray = nil
	pMenu.menu = ""

	return pMenu
}

// StoreMenu stores you menu object into a custom map, indexed by a string identifier
func (menu *Menu) StoreMenu(identifier string) {
	menuDataStore[identifier] = menu
}

// DisplayStoredMenu displays a Menu form the menuStore base on the identifier
func DisplayStoredMenu(identifier string) *Menu {
	return menuDataStore[identifier]
}

// DisplayMenu displays the menu.
func (menu *Menu) DisplayMenu() {
	menu.drawField()
	stringArray := make([]string, menu.height)

	for i := 0; i < len(menu.menuArray); i++ {
		stringArray[i] = strings.Join(menu.menuArray[i], "")
	}

	menu.menu = strings.Join(stringArray, "\n")

	fmt.Println(menu.menu)
}
