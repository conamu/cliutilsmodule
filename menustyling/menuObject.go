package menustyling

// Menu is a type struct that holds data to build a menu.
type Menu struct {
	MenuText       [][]string
	StyleChar      string
	LineThicness   int
	SeparatorLines bool
}

// MenuStore gets used to store multiple Menus with string identifiers
// to use them easily
type MenuStore map[string]Menu
