package models

type Menu struct {
	Elements []string
}

type PageHeader struct {
	Title       string
	Description string
	HeaderTitle string
	Menu        Menu
}

func (*Menu) Generate() Menu {

	var menu Menu

	menu.Elements[0] = "Home"

	return menu
}
