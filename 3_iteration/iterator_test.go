package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat 5 times", func(t *testing.T) {

		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}

	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	// Output: aaaaaaaaaa
}
