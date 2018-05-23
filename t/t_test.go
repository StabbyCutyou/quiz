package t

import (
	"testing"

	"github.com/StabbyCutyou/quiz"
)

func TestT(t *testing.T) {
	s := quiz.S{}
	// Not imported because this isn't a _test package matching the regular package
	// and it follows different rules
	//s.MagicUnexported() // s.MagicUnexported undefined (type quiz.S has no field or method MagicUnexported)
	s.Exported()
}
