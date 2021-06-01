package main

import (
	"reflect"
	"testing"
)

func TestWalk_V1(t *testing.T) {

	// Test for a function `walk` which takes any struct and a method, and calls
	// all the method passing in all the `string` fields of the struct.

	expected := "Chris"
	var got []string // slice of strings

	x := struct { // anonymous struct
		Name string
	}{expected}

	walk_v1(x, func(input string) { // anonymous function
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}

func TestWalk_V2(t *testing.T) {

	// Test for a function `walk` which takes any struct and a method, and calls
	// all the method passing in all the `string` fields of the struct.

	expected := "Chris"
	var got []string // slice of strings

	x := struct { // anonymous struct
		Name string
	}{expected}

	walk_v2(x, func(input string) { // anonymous function
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
}

func TestWalk_v3(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk_v3(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
