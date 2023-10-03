package factoradic

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {

	var tests = []struct {
		base10        int32
		factorialBase string
	}{
		{base10: -1, factorialBase: "###"},
		{base10: 0, factorialBase: "0"},
		{base10: 1, factorialBase: "1"},
		{base10: 2, factorialBase: "10"},
		{base10: 3, factorialBase: "11"},
		{base10: 4, factorialBase: "20"},
		{base10: 5, factorialBase: "21"},
		{base10: 6, factorialBase: "100"},
		{base10: 7, factorialBase: "101"},
		{base10: 21, factorialBase: "311"},
		{base10: 22, factorialBase: "320"},
		{base10: 23, factorialBase: "321"},
		{base10: 24, factorialBase: "1000"},
		{base10: 25, factorialBase: "1001"},
		{base10: 5038, factorialBase: "654320"},
		{base10: 5039, factorialBase: "654321"},
		{base10: 5040, factorialBase: "1000000"},
		{base10: 999998, factorialBase: "266251210"},
		{base10: 999999, factorialBase: "266251211"},
		{base10: 1000000, factorialBase: "266251220"},
		{base10: 1000001, factorialBase: "266251221"},
		{base10: 3628799, factorialBase: "987654321"},
		{base10: 3628800, factorialBase: "###"},
	}

	for _, test := range tests {
		want := test.factorialBase
		got := Number(test.base10).String()
		if want != got {
			t.Errorf("(%d) want %s, got %s", test.base10, want, got)
		}
	}
}

func TestParseNumber(t *testing.T) {

	var tests = []struct {
		base10        int32
		factorialBase string
		err           bool
	}{
		{factorialBase: "0", base10: 0},
		{factorialBase: "1", base10: 1},
		{factorialBase: "10", base10: 2},
		{factorialBase: "11", base10: 3},
		{factorialBase: "20", base10: 4},
		{factorialBase: "21", base10: 5},
		{factorialBase: "100", base10: 6},
		{factorialBase: "101", base10: 7},
		{factorialBase: "311", base10: 21},
		{factorialBase: "320", base10: 22},
		{factorialBase: "321", base10: 23},
		{factorialBase: "1000", base10: 24},
		{factorialBase: "1001", base10: 25},
		{factorialBase: "654320", base10: 5038},
		{factorialBase: "654321", base10: 5039},
		{factorialBase: "1000000", base10: 5040},
		{factorialBase: "266251210", base10: 999998},
		{factorialBase: "266251211", base10: 999999},
		{factorialBase: "266251220", base10: 1000000},
		{factorialBase: "266251221", base10: 1000001},
		{factorialBase: "987654321", base10: 3628799},
		{factorialBase: "", err: true},
		{factorialBase: "2", err: true},
		{factorialBase: "102", err: true},
		{factorialBase: "Z", err: true},
		{factorialBase: "1111111111", err: true},
		{factorialBase: "30", err: true},
		{factorialBase: "400", err: true},
	}

	for _, test := range tests {
		want := Number(test.base10)
		got, err := ParseNumber(test.factorialBase)
		if !test.err && err != nil {
			t.Errorf("(%s) unexpected error %v", test.factorialBase, err)
			continue
		}
		if test.err {
			if err == nil {
				t.Errorf("(%s) expected error, got base 10 value %d", test.factorialBase, got)
			}
			continue
		}
		if want != got {
			t.Errorf("(%s) want base 10 value %d, got %d", test.factorialBase, want, got)
		}
	}
}

func ExampleNumber() {
	var n Number = 23
	fmt.Printf("Your number is %s.", n)
	// Output:
	// Your number is 321.
}

func ExampleParseNumber() {
	n, _ := ParseNumber("321")
	fmt.Printf("Your number in base 10 is %d.", n)
	// Output:
	// Your number in base 10 is 23.
}
