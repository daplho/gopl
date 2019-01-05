package methods_test

import (
	"strings"
	"testing"
	"time"

	"github.com/daplho/gopl/ch12/methods"
)

func ExamplePrintDuration() {
	methods.Print(time.Hour)
}

func ExamplePrintReplacer() {
	methods.Print(new(strings.Replacer))
}

func TestExampleOutput(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrintDuration()
	ExamplePrintReplacer()
}
