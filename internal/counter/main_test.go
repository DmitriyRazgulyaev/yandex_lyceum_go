package counter

import "testing"

func TestIncrement(t *testing.T) {
	Increment()

	if counter != 1 {
		t.Error("counter is expected to be incremented")
	}
}

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}
