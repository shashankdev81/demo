package mathlib

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestMultiplyHappyPath(t *testing.T) {
	t.Parallel()
	actual := Multiply(2, 3)
	expected := 6
	if expected != actual {
		t.Errorf("Exepectd %d but got %d ", expected, actual)
	}
}

func TestDivideHappyPath(t *testing.T) {
	t.Parallel()
	actual := Divide(6, 3)
	expected := 2
	if expected != actual {
		t.Errorf("Exepectd %d but got %d ", expected, actual)
	}
}

func TestDivideFailure(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Exepectd exception but code did not panic")

		}
	}()
	Divide(2, 0)
}

type testCase struct {
	arg1     int
	arg2     int
	expected int
	err      bool
}

func TestMultiplySuite(t *testing.T) {
	t.Parallel()
	cases := []testCase{
		{2, 3, 6, false},
		{2, 0, 0, false},
		{2, 1, 2, false},
		{3, 2, 6, false},
		{0, 2, 0, false},
		{1, 2, 2, false},
	}

	for _, tc := range cases {
		actual := Multiply(tc.arg1, tc.arg2)
		if actual != tc.expected {
			t.Errorf("Expected %d and actual %d dont match", tc.expected, actual)
		}
	}
}

func TestDivideSuite(t *testing.T) {
	t.Parallel()
	cases := []testCase{
		{6, 3, 2, false},
		{0, 2, 0, false},
		{2, 2, 1, false},
		{2, 0, 1, true},
		{0, 0, 1, true},
	}

	for _, tc := range cases {
		if tc.err {
			continue
		}
		t.Run(fmt.Sprintf("Success sceanrio | Input a=%v, b=%v, expected=%v", tc.arg1, tc.arg2, tc.expected), func(t *testing.T) {
			actual := Divide(tc.arg1, tc.arg2)
			if actual != tc.expected {
				t.Errorf("Expected %d and actual %d dont match", tc.expected, actual)
			}
		})
	}

	for _, tc := range cases {
		if !tc.err {
			continue
		}
		t.Run(fmt.Sprintf("Failure sceanrio | Input a=%v, b=%v", tc.arg1, tc.arg2), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Exepectd exception but code did not panic")

				}
			}()
			Divide(tc.arg1, tc.arg2)
		})
	}
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	text := "hello world"
	Print(text, &buf)
	if text != strings.TrimSpace(buf.String()) {
		t.Errorf("expected %v but got %v", text, &buf)
	}
}
