package passw

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	charsSet := make(map[byte]struct{})
	for _, c := range chars {
		charsSet[c] = struct{}{}
	}

	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			out, err := Generate()
			Ok(t, err)
			if testing.Verbose() {
				t.Logf("%s", out)
			}
			Equal(t, 15, len(out), "incorrect output length") // (4 * 3) + 3

			for j := 0; j < len(out); j++ {
				if j == 3 || j == 7 || j == 11 {
					Equal(t, byte('-'), out[j], "invalid byte %v at index [%d]", out[j], j)
				} else if _, ok := charsSet[out[j]]; !ok {
					t.Errorf("invalid byte %v at index [%d]", out[j], j)
				}
			}
		})
	}
}

func Ok(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func Equal(t *testing.T, expect, got interface{}, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if expect != got {
		t.Errorf("%s: expected: %v, got: %v", msg, expect, got)
	}
}
