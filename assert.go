package assert

import (
	"strings"
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

func Empty(t *testing.T, s string) {
	if strings.TrimSpace(s) != "" {
		t.Fatalf("string was not empty: %v", s)
	}
}

func NotEmpty(t *testing.T, s string) {
	if strings.TrimSpace(s) == "" {
		t.Fatalf("string was empty")
	}
}

func of[T comparable](s []T, has ...T) bool {
	source := make(map[T]struct{}, 0)
	matched := make(map[T]struct{}, 0)

	for _, v := range has {
		source[v] = struct{}{}
	}

	for _, a := range s {
		if len(has) < 1 {
			break
		}
		for b := range source {
			if a == b {
				matched[b] = struct{}{}
				break
			}
		}
	}
	return len(matched) == len(source)
}

func Of[T comparable](t *testing.T, slice []T, has ...T) {
	if !of(slice, has...) {
		t.Fatalf("%#v did not contain %#v", slice, has)
	}
}

func NotOf[T comparable](t *testing.T, slice []T, has ...T) {
	if of(slice, has...) {
		t.Fatalf("%#v did not contain %#v", slice, has)
	}
}

func in[K, V comparable](m map[K]V, has map[K]V) bool {
	matched := make(map[K]V, 0)

	for k, v := range m {
		if v == has[k] {
			matched[k] = v
		}
	}
	return len(matched) == len(has)
}

func In[K, V comparable](t *testing.T, m map[K]V, has map[K]V) {
	if !in(m, has) {
		t.Fatalf("%#v did not contain %#v", m, has)
	}
}

func NotIn[K, V comparable](t *testing.T, m map[K]V, has map[K]V) {
	if in(m, has) {
		t.Fatalf("%#v contain %#v", m, has)
	}
}
