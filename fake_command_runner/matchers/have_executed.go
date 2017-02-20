package fake_command_runner_matchers // import "code.cloudfoundry.org/commandrunner/fake_command_runner/matchers"

import (
	"fmt"
	"os/exec"

	"code.cloudfoundry.org/commandrunner/fake_command_runner"
)

// HaveExecuted is like HaveExecutedSerially, but the commands can be in any order.
func HaveExecuted(specs ...fake_command_runner.CommandSpec) *HaveExecutedMatcher {
	return &HaveExecutedMatcher{Specs: specs}
}

type HaveExecutedMatcher struct {
	Specs    []fake_command_runner.CommandSpec
	executed []*exec.Cmd
}

func (m *HaveExecutedMatcher) Match(actual interface{}) (bool, error) {
	runner, ok := actual.(*fake_command_runner.FakeCommandRunner)
	if !ok {
		return false, fmt.Errorf("Not a fake command runner: %#v.", actual)
	}

	m.executed = runner.ExecutedCommands()

	matched := false

	for _, spec := range m.Specs {
		matched = false

		for i := 0; i < len(m.executed); i++ {

			if !spec.Matches(m.executed[i]) {
				continue
			}

			matched = true

			break
		}

		if !matched {
			break
		}
	}

	if matched {
		return true, nil
	} else {
		return false, nil
	}
}

func (m *HaveExecutedMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected to execute:%s\n\nActually executed:%s", prettySpecs(m.Specs), prettyCommands(m.executed))
}

func (m *HaveExecutedMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected to not execute the following commands:%s", prettySpecs(m.Specs))
}
