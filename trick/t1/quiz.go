package t1

import "fmt"

// S is a struct
type S struct{}

// Exported is exported
func (s *S) Exported() {
	fmt.Println("Exported!")
}

func (s *S) unexported() {
	fmt.Println("Unexported!")
}
