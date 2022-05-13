package pages

import (
	"github.com/KirillBogdanets/pw-go/po/components"
	pw "github.com/KirillBogdanets/pw-go/pw"
)

type TodosPage struct {
	SearchBar components.SearchBar
	*pw.PwWrapper
}

func NewTodosPage() *TodosPage {
	return &TodosPage{}
}
