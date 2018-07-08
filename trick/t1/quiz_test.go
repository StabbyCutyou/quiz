package t1_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/StabbyCutyou/quiz/trick/t1"
)

func TestQuiz(t *testing.T) {
	s := t1.S{}
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	// Imported because this is a _test package matching the regular package
	// and thus follows different rules
	s.MagicUnexported()
}
