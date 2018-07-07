package names

import (
	"fmt"
	"os"
	"testing"
)

func TestNames(t *testing.T) {
	// ignoredFunc() not imported!
	fmt.Println(os.Args[0])
}
