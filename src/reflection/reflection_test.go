package main

import (
	"reflect"
	"testing"
)

// Only checking for number of calls
func TestWalk_v1(t *testing.T) {

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

// Checking for the actual string values in a struct
func TestWalk_v2(t *testing.T) {

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

// Multiple string fields, plus structs with a non strict field
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
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk_v3(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("%s: got %v, want %v", test.Name, got, test.ExpectedCalls)
			}
		})
	}
}

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

// Nested struct
func TestWalk_v4(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk_v4(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("%s: got %v, want %v", test.Name, got, test.ExpectedCalls)
			}
		})
	}
}

// Nested fields, pointers, slices
func TestWalk_v5(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	functionsToTest := []func(interface{}, func(string)){
		walk_v5,
		walk_v5_r1,
		walk_v5_r2,
	}

	for _, test := range cases {
		for _, funcUnderTest := range functionsToTest {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				funcUnderTest(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("%s: got %v, want %v", test.Name, got, test.ExpectedCalls)
				}
			})
		}
	}
}

// Nested fields, pointers, slices, arrays
func TestWalk_v6(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	functionsToTest := []func(interface{}, func(string)){
		walk_v6,
	}

	for _, test := range cases {
		for _, funcUnderTest := range functionsToTest {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				funcUnderTest(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("%s: got %v, want %v", test.Name, got, test.ExpectedCalls)
				}
			})
		}
	}
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}

// Nested fields, pointers, slices, arrays, maps
func TestWalk_v7(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	functionsToTest := []func(interface{}, func(string)){
		walk_v7,
		walk_v7_r1,
	}

	aMap := map[string]string{
		"Foo": "Bar",
		"Baz": "Boz",
	}

	for _, funcUnderTest := range functionsToTest {
		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				funcUnderTest(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("%s: got %v, want %v", test.Name, got, test.ExpectedCalls)
				}
			})
		}

		t.Run("with maps", func(t *testing.T) {
			var got []string
			funcUnderTest(aMap, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, "Bar")
			assertContains(t, got, "Boz")
		})
	}
}

// Nested fields, pointers, slices, arrays, maps, channels, functions
func TestWalk_v8(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	functionsToTest := []func(interface{}, func(string)){
		walk_v8,
	}

	aMap := map[string]string{
		"Foo": "Bar",
		"Baz": "Boz",
	}

	for _, funcUnderTest := range functionsToTest {
		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				funcUnderTest(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("%s: got %v, want %v", test.Name, got, test.ExpectedCalls)
				}
			})
		}

		t.Run("with maps", func(t *testing.T) {
			var got []string
			funcUnderTest(aMap, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, "Bar")
			assertContains(t, got, "Boz")
		})

		t.Run("with channels", func(t *testing.T) {
			aChannel := make(chan Profile)

			go func() {
				aChannel <- Profile{33, "Berlin"}
				aChannel <- Profile{34, "Katowice"}
				close(aChannel)
			}()

			var got []string
			want := []string{"Berlin", "Katowice"}

			walk_v8(aChannel, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("with function", func(t *testing.T) {
			aFunction := func() (Profile, Profile) {
				return Profile{33, "Berlin"}, Profile{34, "Katowice"}
			}

			var got []string
			want := []string{"Berlin", "Katowice"}

			walk_v8(aFunction, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}

		})
	}
}
