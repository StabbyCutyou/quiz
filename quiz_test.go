package quiz_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/StabbyCutyou/quiz"
)

func TestQuiz(t *testing.T) {
	s := quiz.S{}
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	// Imported because this is a _test package matching the regular package
	// and it follows different rules
	s.MagicUnexported()
}
