package snow

import (
	"bytes"
	"fmt"
)

// Operator is the type to define all operators
type Operator string

// All Operators available
const (
	Equal    Operator = "="
	NotEqual Operator = "!="
	Greater  Operator = ">"
	Less     Operator = "<"
	IN       Operator = "IN"
	Like     Operator = "LIKE"
)

// LogicalOperator is the type to define the logical operators
type LogicalOperator string

// All LogicalOperators Available
const (
	AND LogicalOperator = "^"
	OR LogicalOperator = "^OR"
)

// QueryElement is a interface to deal with all element in the query
type QueryElement interface{
	String()string
}

type LogicalQuery struct{
	LogicalOperator
}

// String to get the string value from the Logical Operator
func (lo *LogicalQuery) String()string{
	return string(lo.LogicalOperator)
}

// FieldQuery struct to define name, value e operator
type FieldQuery struct {
	Name     string
	Value    string
	Operator Operator
}

// String to covert the FieldQuery in string
func (fq *FieldQuery) String() string {
	return fmt.Sprintf("%s%s%s", fq.Name, fq.Operator, fq.Value)
}

// Query store the query elements
type Query []QueryElement

// String to covert the Query in string
func (q *Query) String() string{
	var buffer bytes.Buffer
	for _, item := range *q{
		buffer.WriteString(item.String())
	}
	return buffer.String()
}