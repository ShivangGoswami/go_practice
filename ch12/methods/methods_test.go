package methods_test

import (
	"strings"
	"testing"
	"time"

	"../methods"
)

func Test(t *testing.T) {
	methods.Print(time.Hour)
	methods.Print(new(strings.Replacer))
}
func ExamplePrintDuration() {
	methods.Print(time.Hour)
}

func ExamplePrintReplacer() {
	methods.Print(new(strings.Replacer))
}
