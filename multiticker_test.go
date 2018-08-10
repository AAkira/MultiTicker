package multiticker

import (
	"testing"
	"time"
	"context"
	"fmt"
)

func TestMultiTicker(t *testing.T) {
	t.Helper()

	type testCase struct {
		name          string
		list          map[string]int
		timeoutSecond int
	}

	testCases := []testCase{
		{
			name: "normal",
			list: map[string]int{
				"a": 2,
				"b": 4,
				"c": 6,
				"d": 12,
			},
			timeoutSecond: 9,
		},
		{
			name: "prime number",
			list: map[string]int{
				"a": 3,
				"b": 7,
			},
			timeoutSecond: 11,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			loopCounter := make(map[string]int, len(testCase.list))
			loopCountDefault := 0

			intervalList := make(map[string]time.Duration, len(testCase.list))
			for k, v := range testCase.list {
				intervalList[k] = time.Duration(v) * time.Second
				loopCounter[k] = 0
			}
			ticker := NewMultiTicker(intervalList)
			defer ticker.Stop()

			ctx, cancel := context.WithTimeout(context.Background(),
				time.Duration(testCase.timeoutSecond)*time.Second)
			defer cancel()

			fmt.Println(loopCounter)
		L:
			for {
				select {
				case c := <-ticker.C:
					if _, ok := loopCounter[c.Key]; ok {
						loopCounter[c.Key]++
						fmt.Printf("%s %d", c.Key, loopCounter[c.Key])
					} else {
						loopCountDefault++
					}
				case <-ctx.Done():
					break L
				}
			}

			for k, v := range loopCounter {
				// check caller count
				fmt.Printf("%s %d %d \n", k, v, testCase.timeoutSecond/testCase.list[k])
				if v != testCase.timeoutSecond/testCase.list[k] {
					t.Fail()
				}
			}
			if loopCountDefault != 0 {
				t.Fail()
			}
		})
	}

}

func TestMultiTicker_gcd(t *testing.T) {
	t.Helper()

	type testCase struct {
		name     string
		inputA   int
		inputB   int
		expected int
	}

	testCases := []testCase{
		{
			name:     "1, 2",
			inputA:   1,
			inputB:   2,
			expected: 1,
		},
		{
			name:     "1, 7",
			inputA:   1,
			inputB:   7,
			expected: 1,
		},
		{
			name:     "2, 4",
			inputA:   2,
			inputB:   4,
			expected: 2,
		},
		{
			name:     "prime number",
			inputA:   7,
			inputB:   11,
			expected: 1,
		},
		{
			name:     "0, 1",
			inputA:   0,
			inputB:   1,
			expected: 1,
		},
		{
			name:     "10, 5",
			inputA:   10,
			inputB:   5,
			expected: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if gcd(testCase.inputA, testCase.inputB) != testCase.expected {
				t.Fail()
			}
		})
	}
}

func TestMultiTicker_max(t *testing.T) {
	t.Helper()

	type testCase struct {
		name     string
		inputA   int
		inputB   int
		expected int
	}

	testCases := []testCase{
		{
			name:     "1, 2",
			inputA:   1,
			inputB:   2,
			expected: 2,
		},
		{
			name:     "0, 1",
			inputA:   0,
			inputB:   1,
			expected: 1,
		},
		{
			name:     "10, 4",
			inputA:   10,
			inputB:   4,
			expected: 10,
		},
		{
			name:     "11, 7",
			inputA:   11,
			inputB:   7,
			expected: 11,
		},
		{
			name:     "10, 10",
			inputA:   10,
			inputB:   10,
			expected: 10,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if max(testCase.inputA, testCase.inputB) != testCase.expected {
				t.Fail()
			}
		})
	}
}
