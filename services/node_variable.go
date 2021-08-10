package services

import (
	"fmt"
	"strings"
)

// NodeVariable represents an attribute of a service
// in scenario
// NodeName is the node's unique name or id
// AttributeChain is an array that shows the field of that
// node. For example, [a, b, c] means $variable.a.b.c.
type NodeVariable struct {
	NodeName       string
	AttributeChain []string
}

// String returns string representation of a NodeVariable.
func (v *NodeVariable) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("$%s", v.NodeName))
	for _, attr := range v.AttributeChain {
		sb.WriteString(fmt.Sprintf(".%s", attr))
	}

	return sb.String()
}
