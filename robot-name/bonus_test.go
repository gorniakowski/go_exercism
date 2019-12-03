// +build bonus

package robotname

import (
	"fmt"
	"testing"
)

var maxNames = 26 * 26 * 10 * 10 * 10

func TestCollisions(t *testing.T) {
	// Test uniqueness for new robots.
	for i := len(seen); i <= maxNames-600000; i++ {
		fmt.Println("new:", len(seen))
		_ = New().getName(t, false)
	}

	// Test that names aren't recycled either.
	r := New()
	for i := len(seen); i < maxNames; i++ {
		fmt.Println(len(seen))
		r.Reset()
		_ = r.getName(t, false)
	}

	// Test that name exhaustion is handled more or less correctly.
	_, err := New().Name()
	fmt.Println(len(seen))
	if err == nil {
		t.Fatalf("should return error if namespace is exhausted")
	}
}
