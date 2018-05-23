package quiz

func (s *S) MagicUnexported() {
	s.unexported()
}
