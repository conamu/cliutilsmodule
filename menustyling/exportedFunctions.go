package menustyling

import (
	"fmt"
	"github.com/conamu/cliutilsmodule/utils"
	"strings"
)

// menuDataStore is a Map which stores a max of 20 menus to use later
var menuDataStore = make(menuStore, 20)
var tableDataStore = make(tableStore, 20)

// CreateMenu Creates and Initialises a new Menu fo you to Customize.
func CreateMenu(menuText [][]string, styleChar string, lineThickness int, whiteSpace int, separatorLines bool, takeInput bool) *Menu {
	var menu Menu
	var pMenu = &menu

	pMenu.height = 0
	pMenu.width = 0
	pMenu.stringArray = nil
	pMenu.menuArray = nil
	pMenu.menu = ""
	pMenu.lineThickness = lineThickness
	pMenu.separatorLines = separatorLines
	pMenu.styleChar = styleChar
	pMenu.menuText = menuText
	pMenu.whiteSpace = whiteSpace
	pMenu.takeInput = takeInput

	return pMenu
}

// CreateTable Creates and Initialises a new Table fo you to Customize.
func CreateTable(tableText [][]string, styleChar string, lineThickness int) *Table {
	var table Table
	var pTable = &table

	pTable.height = 0
	pTable.width = 0
	pTable.lineThickness = lineThickness
	pTable.tableText = tableText
	pTable.styleChar = styleChar
	pTable.tableArray = nil
	pTable.stringArray = nil

	return pTable
}

// StoreTable stores your Table object into a custom map, indexed by a string identifier for quick retrieval.
func (table *Table) StoreTable(identifier string) {
	tableDataStore[identifier] = table
}

// StoreMenu stores your Menu object into a custom map, indexed by a string identifier for quick retrieval.
func (menu *Menu) StoreMenu(identifier string) {
	menuDataStore[identifier] = menu
}

// GetStoredMenu returns a Menu form the menuStore based on the identifier
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
	fmt.Println(menu.menu)
	if menu.takeInput {
		menu.inputData = utils.GetIO()
	}
}

func (menu *Menu) GetInputData() string {
	return menu.inputData
}

