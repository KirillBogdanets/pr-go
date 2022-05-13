package components

import (
	"github.com/KirillBogdanets/pw-go/helpers"
	pw "github.com/KirillBogdanets/pw-go/pw"
	"github.com/playwright-community/playwright-go"
)

var customAssertions helpers.CustomAssertions

var (
	searchField          = "input.new-todo"
	searchResults        = "ul.todo-list > li"
	inputToggle          = "input.toggle"
	activeTab            = "text=Active"
	completedTab         = "text=Completed"
	clearCompletedButton = "text=Clear completed"
)

type SearchBar struct {
	PwWrapper pw.PwWrapper
}

func (s *SearchBar) Search(initializedPage playwright.Page, text string) {
	s.PwWrapper.Click(initializedPage, searchField)
	s.PwWrapper.TypeText(initializedPage, searchField, text)
	s.PwWrapper.Press(initializedPage, searchField, "Enter")
}

func (s *SearchBar) GetSearchResultsLength(initializedPage playwright.Page) int {
	arrayOfElements, err := s.PwWrapper.GetListOfElements(initializedPage, searchResults)
	customAssertions.AssertErrorToNilf("could not get array of elements: %w", err)
	return len(arrayOfElements)
}

func (s *SearchBar) SelectInputToggle(initializedPage playwright.Page) {
	s.PwWrapper.Click(initializedPage, inputToggle)
}

func (s *SearchBar) SelectActiveTab(initializedPage playwright.Page) {
	s.PwWrapper.Click(initializedPage, activeTab)
}

func (s *SearchBar) SelectCompletedTab(initializedPage playwright.Page) {
	s.PwWrapper.Click(initializedPage, completedTab)
}

func (s *SearchBar) ClearCompleted(initializedPage playwright.Page) {
	s.PwWrapper.Click(initializedPage, clearCompletedButton)
}
