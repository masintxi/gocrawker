package main

import (
	"reflect"
	testing "testing"
)

func TestPrintReport(t *testing.T) {
	test := map[string]struct {
		input    map[string]int
		expected []Page
	}{
		"first": {
			input: map[string]int{
				"url1": 1,
				"url2": 2,
				"url3": 3,
			},
			expected: []Page{
				{"url3", 3},
				{"url2", 2},
				{"url1", 1},
			},
		},
		"second": {
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			expected: []Page{
				{"d", 4},
				{"c", 3},
				{"b", 2},
				{"a", 1},
			},
		},
		"third": {
			input: map[string]int{
				"d": 1,
				"c": 1,
				"b": 1,
				"a": 1,
				"e": 2,
			},
			expected: []Page{
				{"e", 2},
				{"a", 1},
				{"b", 1},
				{"c", 1},
				{"d", 1},
			},
		},
		"empty": {
			input:    map[string]int{},
			expected: []Page{},
		},
		"nil": {
			input:    nil,
			expected: []Page{},
		},
		"one item": {
			input: map[string]int{
				"a": 1,
			},
			expected: []Page{
				{"a", 1},
			},
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual := sortPages(tc.input)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - FAIL: expected %v, got %v", name, tc.expected, actual)
				return
			}
		})
	}
}
