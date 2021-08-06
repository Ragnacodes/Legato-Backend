package services

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

// Regular expressions used for variables
const (
	NodeVariableRegex          = "\\$([a-zA-Z0-9]*)(\\.([a-zA-Z0-9]*))*"
	NodeVariableNameRegEx      = "\\$([a-zA-Z0-9]*)"
	NodeVariableAttributeRegEx = "(\\.([a-zA-Z0-9]*))"
)

// parseNodeVariable returns the NodeVariable of a string variable
func parseNodeVariable(v string) NodeVariable {
	// Compile some regex rules to parse the string
	rgForName, _ := regexp.Compile(NodeVariableNameRegEx)
	rgForAttributes, _ := regexp.Compile(NodeVariableAttributeRegEx)

	// Parse the variable name
	variableName := rgForName.FindString(v)[1:]
	log.Println("name:", variableName)

	// Parse the attributes of the variable
	attributesWithDots := rgForAttributes.FindAllString(v, -1)
	var attributes []string
	attributes = []string{}
	for _, att := range attributesWithDots {
		attributes = append(attributes, att[1:])
	}
	log.Println("attributes: ", attributes)

	return NodeVariable{
		NodeName:       variableName,
		AttributeChain: attributes,
	}
}

// Parse looks up the input string and will find all of the
// values for variables and replace them in the input string.
// At last it will return output string with no variables.
func Parse(input string) (string, error) {
	// Compile some regex rules to parse the string
	rg, _ := regexp.Compile(NodeVariableRegex)

	// Find all of variables in the text.
	// TODO: what is exactly the n argument of FindAddString()
	stringVars := rg.FindAllString(input, -1)
	log.Println("stringVars:", stringVars)

	// Create NodeVariable object for each variable
	// This contains parsing the string Variables' name and attributes
	var nodeVariables []NodeVariable
	for _, variable := range stringVars {
		nodeVariables = append(nodeVariables, parseNodeVariable(variable))
	}

	// Searching for NodeVariables' real value
	// In the next step we're gonna replace these values
	var vars = map[string]string{}
	for _, nr := range nodeVariables {
		// Find value of this variable in scenario
		v, err := searchForVariable(nr)
		if err != nil {
			log.Println("Error:", err)
		}

		// Add the value to the map
		vars[nr.String()] = v
	}
	// Print all of the variables
	for _, nr := range nodeVariables {
		log.Println(nr.String(), "is equal to", vars[nr.String()])
	}

	// Replace all of the variables with their real values
	realValue := input
	for _, nr := range nodeVariables {
		realValue = strings.ReplaceAll(realValue, nr.String(), vars[nr.String()])
	}

	return realValue, nil
}

// searchForVariable is a temporary function that returns an arbitrary value
func searchForVariable(variable NodeVariable) (string, error) {
	r := rand.Intn(30)

	// Sometimes an error might be occurred
	if r > 28 {
		return "", errors.New("there is not any node with this name")
	}

	return fmt.Sprintf("%d", r), nil
}
