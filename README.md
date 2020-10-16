# Cli-Utils-Module @v1.1.1

## Menustyling functionality

#### With this module you can create nice-looking menus for your CLI apps. They are text-based.

#### `CreateMenu()` will take your textblocks and some parameters. it will return a pointer to an object with the Type Menu.

#### First, create a 2-dimensional array for your menu text-blocks. Example: 
```
menuText := make([][]string, 2)`
menuText[0] = []string{"Main Menu", "Please choose an option below."}`
menuText[1] = []string{"1) Yes", "2) No"}`
```
#### Then, you configure the menu with these arguments:
1. 2 dimensional string array: your `menuText` arrray
2. string: the styling character for the menu. f.E.: `"@"`
3. int: the line thickness of the menu boarders
4. int: the ammount of whitespace between text and boarders
4. bool: do you want separation lines between your text blocks?
5. bool: do you want to take input after displaying the menu?

#### Example: `CreateMenu(menuText, "=", 2, 5, true, true)`

#### After this you can:
1. Display your menu  
`CreateMenu(menuText, "=", 2, 5, true, true).DisplayMenu()`
2. store it in a variable with type Menu  
`main := CreateMenu(menuText, "=", 2, 5, true, true)`
3. store it inside the Menustore with a string ID  
`CreateMenu(menuText, "=", 2, 5, true, true).StoreMenu("main")`

#### You can also get a stored menu and display it from the storage directly
`GetStoredMenu("main").DisplayMenu()`

#### If you got input from a Menu, its stored inside the Menu object. You can access it like that:
`GetStoredMenu("main").GetInputData()`
this will return a **string** for you to use.

### Upcoming features:
1. modifiy a portion of the menutext to simplify templating
2. Animated decoration characters
3. Template menus
4. Option to create Tables instead of menus
5. Option to nest Menus for different styles.
6. Option to have a sliding text if text is too big
7. option to have a max menu size.

## Utility Functions:

### `GetIO()`
#### returns a string after getting input with the bufio scanner

### `GetRandom(max int)`
#### this will return a random int each time. It is seeded with nanoseconds of system time. You hace to specify the max range of numbers you want.

