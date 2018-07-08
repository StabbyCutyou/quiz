package siblings

import "fmt"

// Sibling is a brother, sister, or other individual who shares some notion of
// common parentage or in certain cases an extreme bond with you
type Sibling struct {
	Name string
}

// Fight has no return value, because nothing ever comes from fighting
func (s *Sibling) Fight() {

}

// SayHi will have the package announce itself
func SayHi() {
	fmt.Println("We are the three Siblings of Go Testing!")
}
