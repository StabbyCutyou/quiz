package quiz_test

import (
	"testing"

	"github.com/StabbyCutyou/quiz"
)

func TestQuiz(t *testing.T) {
	s := quiz.S{}
	// Imported because this is a _test package matching the regular package
	// and it follows different rules
	s.MagicUnexported()
}
