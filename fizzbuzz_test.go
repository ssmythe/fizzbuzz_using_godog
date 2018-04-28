package main

import (
	"github.com/DATA-DOG/godog"
	"strconv"
	"fmt"
)

var number int
var result string

func theNumber(arg1 string) error {
	var err error

	number, err = strconv.Atoi(arg1)
	if err != nil {
		return fmt.Errorf("unable to convert number \"%s\", but got \"%s\"", arg1)
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

