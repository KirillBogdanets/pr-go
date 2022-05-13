package tests

import (
	"os"
	"testing"

	"github.com/KirillBogdanets/pw-go/helpers"
	"github.com/KirillBogdanets/pw-go/po/pages"
	"github.com/dailymotion/allure-go"
	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

var utils helpers.Utils
var customAssertions helpers.CustomAssertions
var todosPage = new(pages.TodosPage)
var initializedPage playwright.Page

func SetInitializedPage(initP playwright.Page) {
	initializedPage = initP
}

const TODO_NAME = "Bake a cake"

func BeforeEach(t *testing.T) {
	godotenv.Load()
	todosPage := new(pages.TodosPage)
	initializedPage := todosPage.PwInitialize()
	SetInitializedPage(initializedPage)
	todosPage.Open(initializedPage, "http://todomvc.com/examples/react/")
	todosPage.SearchBar.Search(initializedPage, TODO_NAME)
}

func AfterEach(t *testing.T) {
	todosPage.CloseBrowser(initializedPage)
}

// test function
func TestSearchBarWithOneResult(t *testing.T) {
	BeforeEach(t)
	allure.Test(t,
		allure.Description("Search Bar Test With Only 1 Search Result Item"),
		allure.Parameter("Test Env", os.Getenv("ENV")),
		allure.Action(func() {
			arrayLength := todosPage.SearchBar.GetSearchResultsLength(initializedPage)
			customAssertions.AssertEqual(arrayLength, 1, initializedPage, t)
		}),
	)
	AfterEach(t)
}

func TestSearchBarWithOneResultAfterReload(t *testing.T) {
	BeforeEach(t)
	allure.Test(t,
		allure.Description("Search Bar Test With Only 1 Search Result Item After Reload"),
		allure.Parameter("Test Env", os.Getenv("ENV")),
		allure.Action(func() {
			todosPage.Reload(initializedPage)
			arrayLength := todosPage.SearchBar.GetSearchResultsLength(initializedPage)
			customAssertions.AssertEqual(arrayLength, 1, initializedPage, t)
		}),
	)
	AfterEach(t)
}

func TestSearchBarSmoke(t *testing.T) {
	BeforeEach(t)
	allure.Test(t,
		allure.Description("Search Bar Test With Only 1 Search Result Item After Reload"),
		allure.Parameter("Test Env", os.Getenv("ENV")),
		allure.Action(func() {
			todosPage.SearchBar.SelectInputToggle(initializedPage)

			todosPage.SearchBar.SelectActiveTab(initializedPage)
			arrayLength := todosPage.SearchBar.GetSearchResultsLength(initializedPage)
			customAssertions.AssertEqual(arrayLength, 0, initializedPage, t)

			todosPage.SearchBar.SelectCompletedTab(initializedPage)
			arrayCompletedItemsLength := todosPage.SearchBar.GetSearchResultsLength(initializedPage)
			customAssertions.AssertEqual(arrayCompletedItemsLength, 1, initializedPage, t)

			todosPage.SearchBar.ClearCompleted(initializedPage)
			arrayEmptyLength := todosPage.SearchBar.GetSearchResultsLength(initializedPage)
			customAssertions.AssertEqual(arrayEmptyLength, 0, initializedPage, t)
		}),
	)
	AfterEach(t)
}
