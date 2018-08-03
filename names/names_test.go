package names

import (
	"fmt"
	"os"
	"testing"
)

func TestNames(t *testing.T) {
	// If you try to call ignoredFunc(), you'll see this:
	// ignoredFunc() not imported!
	fmt.Println(os.Args[0])
}
