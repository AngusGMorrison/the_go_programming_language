// Write tests for the charcount program in Section 4.3.
package charcount

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		input       string
		wantCounts  map[rune]int
		wantUTFLens [arrsize]int
		wantInvalid int
	}{
		{
			input:       "",
			wantCounts:  map[rune]int{},
			wantUTFLens: [arrsize]int{},
			wantInvalid: 0,
		},
		{
			input: "hello",
			wantCounts: map[rune]int{
				'h': 1,
				'e': 1,
				'l': 2,
				'o': 1,
			},
			wantUTFLens: [arrsize]int{0, 5},
			wantInvalid: 0,
		},
		{
			input: "heLlo",
			wantCounts: map[rune]int{
				'h': 1,
				'e': 1,
				'l': 1,
				'L': 1,
				'o': 1,
			},
			wantUTFLens: [arrsize]int{0, 5},
			wantInvalid: 0,
		},
		{
			input: "he llo!}",
			wantCounts: map[rune]int{
				'h': 1,
				'e': 1,
				' ': 1,
				'l': 2,
				'o': 1,
				'!': 1,
				'}': 1,
			},
			wantUTFLens: [arrsize]int{0, 8},
			wantInvalid: 0,
		},
		{
			input:       string([]byte{128}), // 127 is highest single-byte UTF-8 value
			wantCounts:  map[rune]int{},
			wantUTFLens: [arrsize]int{},
			wantInvalid: 1,
		},
		{
			input:       "�", // physical 3-byte replacement char should be considered valid input
			wantCounts:  map[rune]int{'�': 1},
			wantUTFLens: [arrsize]int{0, 0, 0, 1},
			wantInvalid: 0,
		},
		{
			input: "a£�𐄂", // 1, 2, 3 and 4-byte UTF-8 characters
			wantCounts: map[rune]int{
				'a': 1,
				'£': 1,
				'�': 1,
				'𐄂': 1,
			},
			wantUTFLens: [arrsize]int{0, 1, 1, 1, 1},
			wantInvalid: 0,
		},
	}

	for _, test := range testCases {
		counts, utflens, invalid, err := Count(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("with input %q got err: %v", test.input, err)
			continue
		}

		for k, v := range counts {
			if v != test.wantCounts[k] {
				t.Errorf("with input %q: want counts = %v, got %v",
					test.input, test.wantCounts, counts)
				break
			}
		}

		for i, count := range utflens {
			if count != test.wantUTFLens[i] {
				t.Errorf("with input %q: want utflens = %v, got %v",
					test.input, test.wantUTFLens, utflens)
				break
			}
		}

		if invalid != test.wantInvalid {
			t.Errorf("with input %q: want invalid = %d, got %d",
				test.input, test.wantInvalid, invalid)
		}
	}
}
