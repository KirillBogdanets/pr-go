# Golang + Playwright Testing Frmaework

The framework based on playwright-go v.7 using BDD approach with chai assertions

## Table of Contents

* [Quickstart](#quickstart)
* [CLI](#cli)
* [Environments variables](#environment-variables)
* [Allure](#allure)

### Quickstart
* Download Java from [go.dev](https://go.dev/dl/) (Recommended Version)
* Install Allure CLI Globally node from [npmjs.com](https://www.npmjs.com/package/allure-commandline)
* To run all tests against use `go test ./tests`

### CLI
* Test run is pretty flexible and could be modified using command line parameters:
    * `-run` name of the test case that you want to run from the whole spec file
    * `-v` name of the spec file that will be run 
    
_Example of command line parameters usage:_

`go test ./tests`
`go test ./tests -run TestSearchBarSmoke`
`go test ./tests -v search_test.go`

### Environment variables
* Environment variables that are used in test run:
    * `ALLURE_RESULTS_PATH` where allure-results should be storaged
        * Default value is `./`
    * `ENV` environment where tests will be run
        * Default value is `DEV`

### Allure
* Allure [allure-go v0.7.0] reports is used as a report tool for that TAF:
    * for running allure report you have to run following command:
    `allure generate allure-results --clean && allure open`
