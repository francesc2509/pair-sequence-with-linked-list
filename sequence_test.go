package sequence

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/francesc2509/pair-sequence-with-linked-list/entities"
)

type testConfig struct {
	pairs    []map[string]int
	node     *entities.Node
	expected *entities.Node
	name     string
}

func assertNode(expected *entities.Node, value *entities.Node) bool {
	return expected == value || (expected != nil && value != nil &&
		expected.Start() == value.Start() &&
		expected.End() == value.End() &&
		assertNode(expected.Next(), value.Next()))
}

func TestJoin(t *testing.T) {
	tests := []func() *testConfig{
		func() *testConfig {
			node := entities.NewNode(map[string]int{"start": 1, "end": 10}, nil)

			return &testConfig{
				pairs:    nil,
				node:     node,
				expected: node,
				name:     "Pairs argument is nil",
			}
		},
		func() *testConfig {
			node := entities.NewNode(map[string]int{"start": 1, "end": 10}, nil)

			return &testConfig{
				pairs:    []map[string]int{},
				node:     node,
				expected: node,
				name:     "Pairs argument is empty",
			}
		},
		func() *testConfig {
			pair := map[string]int{"start": 1, "end": 10}
			var node *entities.Node = nil

			return &testConfig{
				pairs:    []map[string]int{pair},
				node:     node,
				expected: entities.NewNode(pair, nil),
				name:     "Node argument is nil",
			}
		},
	}

	for idx, fn := range tests {
		config := fn()
		result := join(config.pairs, config.node)
		expected := config.expected

		var output bytes.Buffer
		output.WriteString(fmt.Sprintf("Test %d '%s' ", idx, config.name))

		if assertNode(expected, result) {
			output.WriteString(fmt.Sprintf("passed: result %v", result))
			t.Logf(output.String())
		} else {
			output.WriteString(fmt.Sprintf(" did not pass: expected %v, but got %v", expected, result))
			t.Errorf(output.String())
		}
	}
}
