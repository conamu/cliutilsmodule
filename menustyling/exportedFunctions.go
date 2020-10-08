package menustyling

import (
	"fmt"
	"strings"
)

// menuDataStore is a Map which stores a max of 20 menus to use later
var menuDataStore = make(menuStore, 20)

// CreateMenu Creates and Initialises a new Menu fo you to Customize.
func CreateMenu() *Menu {
	var menu Menu
	var pMenu = &menu
	pMenu.height = 0
	pMenu.LineThickness = 0
	pMenu.width = 0
	pMenu.SeparatorLines = false
	pMenu.StyleChar = ""
	pMenu.stringArray = nil
	pMenu.MenuText = nil
	pMenu.menuArray = nil
	pMenu.menu = ""

	return pMenu
}

// StoreMenu stores you menu object into a custom map, indexed by a string identifier
func (menu *Menu) StoreMenu(identifier string) {
	menuDataStore[identifier] = menu
}

// GetStoredMenu displays a Menu form the menuStore base on the identifier
func GetStoredMenu(identifier string) *Menu {
	return menuDataStore[identifier]
}

// DisplayMenu displays the menu.
func (menu *Menu) DisplayMenu() {
	menu.drawField()
	menu.putText()
	menu.stringArray = make([]string, menu.height)
	for i := 0; i < len(menu.menuArray); i++ {
		menu.stringArray[i] = strings.Join(menu.menuArray[i], "")
	}
	menu.menu = strings.Join(menu.stringArray, "\n")
	fmt.Print(menu.menu)
}
