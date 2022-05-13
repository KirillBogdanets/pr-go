package pw

import (
	"github.com/KirillBogdanets/pw-go/helpers"
	"github.com/playwright-community/playwright-go"
)

var customAssertions helpers.CustomAssertions

type PwWrapper struct {
	initializedPage playwright.Page
}

func (pwObj *PwWrapper) PwInitialize() playwright.Page {
	var p PwWrapper
	pw, err := playwright.Run()
	customAssertions.AssertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		Args:     []string{"--incognito"},
	})
	customAssertions.AssertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	customAssertions.AssertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	customAssertions.AssertErrorToNilf("could not create page: %w", err)
	p.initializedPage = page

	return p.initializedPage
}

func (pwObj *PwWrapper) Open(initializedPage playwright.Page, url string) {
	_, err := initializedPage.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	customAssertions.AssertErrorToNilf("could not goto: "+url, err)
}

func (pwObj *PwWrapper) Click(initializedPage playwright.Page, locator string) {
	customAssertions.AssertErrorToNilf("could not click: %v", initializedPage.Click(locator))
}

func (pwObj *PwWrapper) TypeText(initializedPage playwright.Page, locator string, text string) {
	customAssertions.AssertErrorToNilf("could not type: %v", initializedPage.Type(locator, text))
}

func (pwObj *PwWrapper) Press(initializedPage playwright.Page, locator string, keyboardButton string) {
	customAssertions.AssertErrorToNilf("could not press: %v", initializedPage.Press(locator, keyboardButton))
}

func (pwObj *PwWrapper) Reload(initializedPage playwright.Page) {
	initializedPage.Reload()
}

func (pwObj *PwWrapper) GetListOfElements(initializedPage playwright.Page, locator string) ([]playwright.ElementHandle, error) {
	return initializedPage.QuerySelectorAll(locator)
}

func (pwObj *PwWrapper) CloseBrowser(initializedPage playwright.Page) {
	initializedPage.Close()
}
