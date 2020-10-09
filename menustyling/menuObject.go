package menustyling

// Menu is a type struct that holds data to build a menu.
type Menu struct {
	menuText       [][]string
	menuArray      [][]string
	stringArray    []string
	styleChar      string
	inputData      string
	menu           string
	whiteSpace     int
	lineThickness  int
	height         int
	width          int
	separatorLines bool
	takeInput      bool
}

// menuStore gets used to store multiple Menus with string identifiers
// to use them easily
type menuStore map[string]*Menu
