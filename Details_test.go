package problem

import (
	"fmt"
	"testing"

	"github.com/ljpx/test"
)

func TestDetailsAttachError(t *testing.T) {
	// Arrange.
	problem := &Details{}

	// Act.
	problem.AttachError(fmt.Errorf("an error"))

	// Assert.
	test.That(t, problem.Error).IsEqualTo("an error")
}

func TestDetailsAttachErrorNil(t *testing.T) {
	// Arrange.
	problem := &Details{
		Error: "an error",
	}

	// Act.
	problem.AttachError(nil)

	// Assert.
	test.That(t, problem.Error).IsEqualTo("")
}
