package assert

import "testing"

func Eq[T comparable](t *testing.T, a, b T) {
	if a == b {
		return
	}
	t.Fatalf("%v != %v", a, b)
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func GT[T Number](t *testing.T, a, b T) {
	if a > b {
		return
	}
	t.Fatalf("%v < %v", a, b)
}

func LT[T Number](t *testing.T, a, b T) {
	if a < b {
		return
	}
	t.Fatalf("%v > %v", a, b)
}

func GE[T Number](t *testing.T, a, b T) {
	if a >= b {
		return
	}
	t.Fatalf("%v < %v", a, b)
}

func LE[T Number](t *testing.T, a, b T) {
	if a <= b {
		return
	}
	t.Fatalf("%v > %v", a, b)
}

func Nil(t *testing.T, v any) {
	if v == nil {
		return
	}
	t.Fatalf("%v was not nil", v)
}

func NotNil(t *testing.T, v any) {
	if v != nil {
		return
	}
	t.Fatalf("%v was nil", v)
}
