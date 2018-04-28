Feature: FizzBuzz

  Write a program that prints the numbers from 1 to 100. But
  for multiples of three print "Fizz" instead of the number and
  for the multiples of five print "Buzz". For numbers which are
  multiples of both three and five print "FizzBuzz".

  Sample output:
  1
  2
  Fizz
  4
  Buzz
  Fizz
  7
  8
  Fizz
  Buzz
  11
  Fizz
  13
  14
  FizzBuzz
  16
  17
  Fizz
  19
  Buzz
  ...etc up to 100

  Scenario: 1
    Given the number "1"
    When I fizzbuzz
    Then I expect "1"

  Scenario: 2
    Given the number "2"
    When I fizzbuzz
    Then I expect "2"

  Scenario: 3
    Given the number "3"
    When I fizzbuzz
    Then I expect "Fizz"

  Scenario: 5
    Given the number "5"
    When I fizzbuzz
    Then I expect "Buzz"

  Scenario: 15
    Given the number "15"
    When I fizzbuzz
    Then I expect "FizzBuzz"
