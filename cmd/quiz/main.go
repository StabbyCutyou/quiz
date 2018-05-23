package main

import (
	"github.com/StabbyCutyou/quiz"
)

func main() {
	s := quiz.S{}
	// Not compiled into the regular binary, _test.go files aren't included at all
	//s.MagicUnexported() // s.MagicUnexported undefined (type quiz.S has no field or method MagicUnexported)
	s.Exported()
}
