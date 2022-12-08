package assert

import (
	"testing"
)

func TestOf(t *testing.T) {
	Of(t, []int{1, 2, 3}, 1, 1, 1)
	Of(t, []int{1, 2, 3}, 1, 2, 3)
	Of(t, []int{1, 1, 1}, 1)
}

func TestNotOf(t *testing.T) {
	NotOf(t, []int{1, 2, 3}, 4)
	NotOf(t, []int{1, 2, 3}, 4, 5, 6)
}

func TestIn(t *testing.T) {
	In(t, map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2})
}

func TestNotIn(t *testing.T) {
	NotIn(t, map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3})
}
