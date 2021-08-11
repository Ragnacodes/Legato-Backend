package services

import (
	"log"
	"regexp"
)

// Regular expressions used for variables
const (
	NodeVariableRegex          = "\\$([a-zA-Z0-9]*)(\\.([a-zA-Z0-9]*))*"
	NodeVariableNameRegEx      = "\\$([a-zA-Z0-9]*)"
	NodeVariableAttributeRegEx = "(\\.([a-zA-Z0-9]*))"
)

// getStringVariables find the variables and returns
// the string type of node variables.
func getStringVariables(input string) []string {
	rg, _ := regexp.Compile(NodeVariableRegex)

	// Find all of variables in the text.
	// TODO: what is exactly the n argument of FindAddString()
	stringVars := rg.FindAllString(input, -1)
	log.Println("stringVars:", stringVars)
	return stringVars
}

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
