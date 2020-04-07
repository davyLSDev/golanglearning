# Intro

This tutorial comes from [tutorial](https://tutorialedge.net/golang/intro-testing-in-go/)

## Some test commands

These are run in the source directory.

* "go test" - to run the tests
* "go test -v" - to run the test with more verbosity
* "go test -cover" - to see what percentage of code is covered by tests (this is an indicator)
* "go test -coverprofile=coverage.out" - generate a coverprofile
* "go tool cover -html=coverage.out" - generate an html page showing what lines of code are covered. Green = covered, Red = not covered.
