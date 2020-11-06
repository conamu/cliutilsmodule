package menustyling

// Menu is a type struct that holds data to build a Menu.
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

// Table is a type struct that holds data to build a Table.
type Table struct {
	tableText [][]string
	tableArray [][]string
	stringArray []string
	styleChar string
	lineThickness int
	height int
	width int
}

// menuStore gets used to store multiple Menus with string identifiers
// to use them easily
type menuStore map[string]*Menu

// tableStore stores multiple Table pointers to use them easily later.
type tableStore map[string]*Table
