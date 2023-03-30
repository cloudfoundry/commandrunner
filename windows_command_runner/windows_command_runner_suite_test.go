package windows_command_runner_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWindowsCommandRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Windows Command Runner Suite")
}
