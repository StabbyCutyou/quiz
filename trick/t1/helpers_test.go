package t1

// MagicUnexported will only be compiled during a test of this package
// and will not be available in any other context
func (s *S) MagicUnexported() {
	s.unexported()
}
