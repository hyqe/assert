package assert

import (
	"testing"
)

func Want[T comparable](t *testing.T, want, got T) {
	if want == got {
		return
	}
	t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
}

func Eq[T comparable](t *testing.T, a, b T) {
	if a == b {
		return
	}
	t.Fatalf("%v was not equal to %v", a, b)
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func GT[T Number](t *testing.T, a, b T) {
	if a > b {
		return
	}
	t.Fatalf("%v was not greater than %v", a, b)
}

func LT[T Number](t *testing.T, a, b T) {
	if a < b {
		return
	}
	t.Fatalf("%v was not less than %v", a, b)
}

func GE[T Number](t *testing.T, a, b T) {
	if a >= b {
		return
	}
	t.Fatalf("%v is not greater than or equal to %v", a, b)
}

func LE[T Number](t *testing.T, a, b T) {
	if a <= b {
		return
	}
	t.Fatalf("%v is not less than or equal to %v", a, b)
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

func True(t *testing.T, b bool) {
	Want(t, true, b)
}

func False(t *testing.T, b bool) {
	Want(t, false, b)
}

func NoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("error was not nil: %v", err)
	}
}

func Err(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("error was nil")
	}
}
