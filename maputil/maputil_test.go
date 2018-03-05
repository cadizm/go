package maputil

import (
	"sort"
	"testing"
)

func TestSortByValue1(t *testing.T) {
	arr := []RuneIntPair{}

	expected := []RuneIntPair{}

	sort.Sort(ByValue(arr))

	if len(arr) != len(expected) {
		t.Fail()
	}
	for i, v := range arr {
		if v != expected[i] {
			t.Fail()
		}
	}
}

func TestSortByValue2(t *testing.T) {
	arr := []RuneIntPair{
		{'a', 2},
	}

	expected := []RuneIntPair{
		{'a', 2},
	}

	sort.Sort(ByValue(arr))

	if len(arr) != len(expected) {
		t.Fail()
	}
	for i, v := range arr {
		if v != expected[i] {
			t.Fail()
		}
	}
}

func TestSortByValue3(t *testing.T) {
	arr := []RuneIntPair{
		{'a', 31},
		{'b', 42},
		{'c', 17},
		{'d', 26},
	}

	expected := []RuneIntPair{
		{'c', 17},
		{'d', 26},
		{'a', 31},
		{'b', 42},
	}

	sort.Sort(ByValue(arr))

	if len(arr) != len(expected) {
		t.Fail()
	}
	for i, v := range arr {
		if v != expected[i] {
			t.Fail()
		}
	}
}

func TestSortByValue4(t *testing.T) {
	arr := []RuneIntPair{
		{'a', 13},
		{'b', 0},
		{'c', -22},
		{'d', 17},
		{'e', -2},
		{'f', -2},
		{'g', -44},
		{'h', 100},
		{'i', 0},
	}

	expected := []RuneIntPair{
		{'g', -44},
		{'c', -22},
		{'e', -2},
		{'f', -2},
		{'b', 0},
		{'i', 0},
		{'a', 13},
		{'d', 17},
		{'h', 100},
	}

	sort.Sort(ByValue(arr))

	if len(arr) != len(expected) {
		t.Fail()
	}
	for i, v := range arr {
		if v != expected[i] {
			t.Errorf("%v != %v", v, expected[i])
		}
	}
}
