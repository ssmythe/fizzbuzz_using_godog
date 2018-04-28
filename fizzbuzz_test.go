package main

import (
	"github.com/DATA-DOG/godog"
	"os"
	"flag"
	"testing"
	"strconv"
	"fmt"
	"github.com/DATA-DOG/godog/colors"
)

// go test support for godog
var opt = godog.Options{Output: colors.Uncolored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}


// step definitions
var number int
var result string

func theNumber(arg1 string) error {
	var err error

	number, err = strconv.Atoi(arg1)
	if err != nil {
		return fmt.Errorf("unable to convert number \"%s\"", arg1)
	}
	return nil
}

func iFizzbuzz() error {
	result = fizzbuzz(number)
	return nil
}

func iExpect(arg1 string) error {
	if result != arg1 {
		return fmt.Errorf("expected \"%s\", but got \"%s\"", arg1, result)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the number "([^"]*)"$`, theNumber)
	s.Step(`^I fizzbuzz$`, iFizzbuzz)
	s.Step(`^I expect "([^"]*)"$`, iExpect)
}
