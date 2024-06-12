package main

import (
	"log"
	"testing"
)

type testData struct {
	input  string
	expect string
}

func TestUnpackStringValie(t *testing.T) {

	cases := []testData{
		{input: "a4bc2d5e", expect: "aaaabccddddde"},
		{input: "abcd", expect: "abcd"},
		{input: "", expect: ""},
		{input: `qwe\4\5`, expect: "qwe45"},
		{input: `qwe\45`, expect: "qwe44444"},
		{input: `qwe\\5`, expect: `qwe\\\\\`},
	}

	for _, testCase := range cases {
		res, err := unpackString(testCase.input)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		if res != testCase.expect {
			t.Fatalf("Incorrect output: EXPECTED - %s,  GOT - %s\n", testCase.expect, res)
		}

		log.Println("Expected:", testCase.expect, "Got:", res)
	}
}
func TestUnpackStringInvalid(t *testing.T) {
	cases := []testData{
		{input: "45", expect: "error"},
	}

	for _, testCase := range cases {
		_, err := unpackString(testCase.input)
		if err != nil {
			log.Println("Correct behavior")
		} else {
			t.Fatalf("Incorrect behavior")
		}
	}
}
