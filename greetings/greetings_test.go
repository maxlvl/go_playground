package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, 
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Max"
	expected := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Max")
	if !expected.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Max") = %q, %v, expected match for %#q, nil`, msg, err, expected)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, expected "", error`, msg, err)
	}
}