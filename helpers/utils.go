package helpers

import (
	"time"

	"github.com/playwright-community/playwright-go"
)

type Utils struct {
}

func (u *Utils) WaitForTimeout(timeout float64) {
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}

func (u *Utils) MakeScreenshot(initializedPage playwright.Page) ([]byte, error) {
	return initializedPage.Screenshot()
}
