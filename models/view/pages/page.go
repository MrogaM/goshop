package models

type Element struct {
	ID   int32
	Name string
	Link string
}
type Menu struct {
	Elements []Element
}

type PageHeader struct {
	Title       string
	Description string
	HeaderTitle string
	Menu        Menu
}

func (m Menu) Generate(e []Element) Menu {

	m.Elements = make([]Element, len(e))
	return m
}
