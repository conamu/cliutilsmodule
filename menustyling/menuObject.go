package menustyling

// Menu is a type struct that holds data to build a menu.
type Menu struct {
	// Exported Fields the user has to set to generate a menu.
	MenuText       [][]string
	WhiteSpace     int
	StyleChar      string
	LineThickness  int
	SeparatorLines bool
	// In code used fields to store data.
	menuArray      [][]string
	stringArray	   []string
	menu           string
	height         int
	width          int
}

// menuStore gets used to store multiple Menus with string identifiers
// to use them easily
type menuStore map[string]*Menu
