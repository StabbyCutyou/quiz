package siblings_test

import (
	"fmt"
	"testing"

	"github.com/StabbyCutyou/quiz/siblings"
)

func TestFunc(t *testing.T) {

}

func BenchmarkFunc(b *testing.B) {

}

// This will be compiled, but not ran!
func ExampleSayHi() {
	siblings.SayHi()
}

// This will be compiled, and ran, passing because outputs match!
func ExampleSayHi_second() {
	siblings.SayHi()
	// Output: We are the three Siblings of Go Testing!
}

// This will be compiled, and ran, failing because outputs don't match!
func xExampleSayHi_third() {
	siblings.SayHi()
	// Output: We are the four Siblings of Go Testing!
}

func ExampleSibling_Name() {
	s := siblings.Sibling{Name: "Benchy"}
	fmt.Println(s.Name)
	// Output: Benchy
}

func ExampleSibling_Fight() {
	s := siblings.Sibling{Name: "Examy"}
	s.Fight()
	// Output: Fighting isnt very nice
}
