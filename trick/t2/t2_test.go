package t2

import (
	"testing"

	"github.com/StabbyCutyou/quiz/trick/t1"
)

func TestT(t *testing.T) {
	s := t1.S{}
	// Not imported because this isn't a _test package matching the regular package
	// and it follows different rules
	//s.MagicUnexported() // s.MagicUnexported undefined (type quiz.S has no field or method MagicUnexported)
	s.Exported()
}
