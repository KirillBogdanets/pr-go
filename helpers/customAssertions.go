package helpers

import (
	"fmt"
	"log"
	"reflect"

	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/playwright-community/playwright-go"
)

type CustomAssertions struct {
	*Utils
}

func (cA *CustomAssertions) AssertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func (cA *CustomAssertions) AssertEqual(expected, actual interface{}, initializedPage playwright.Page, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		data, err := cA.Utils.MakeScreenshot(initializedPage)
		cA.AssertErrorToNilf("error while Screenshot making", err)
		allure.AddAttachment("Screenshot", allure.ImagePng, data)

		panic(fmt.Sprintf("%v does not equal %v", actual, expected))
	}
}
