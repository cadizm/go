package maputil

import "fmt"

type RuneIntPair struct {
	K rune
	V int
}

func (pair RuneIntPair) String() string {
	return fmt.Sprintf("key: %c, value: %v", pair.K, pair.V)
}

type ByValue []RuneIntPair

func (arr ByValue) Len() int {
	return len(arr)
}

func (arr ByValue) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr ByValue) Less(i, j int) bool {
	return arr[i].V < arr[j].V
}
