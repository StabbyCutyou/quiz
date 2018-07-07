package names_test

import (
	"fmt"
	"os"
	"testing"
)

func ignoredFunc() {}

func TestNames(t *testing.T) {
	fmt.Println(os.Args[0])
}
