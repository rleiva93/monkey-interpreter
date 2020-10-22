// object/object.go

package object

import (
	"bytes"
	"fmt"
	"strings"

	"../ast"
)

// Type string
type Type string

const (
	// INTEGER_OBJ const
	INTEGER_OBJ = "INTEGER"
	// BOOLEAN_OBJ const
	BOOLEAN_OBJ = "BOOLEAN"
	// NULL_OBJ const
	NULL_OBJ = "NULL"
	// RETURN_VALUE_OBJ const
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	// ERROR_OBJ const
	ERROR_OBJ = "ERROR"
	// FUNCTION_OBJ const
	FUNCTION_OBJ = "FUNCTION"
)

// Object interface
type Object interface {
	Type() Type
	Inspect() string
}

// Integer struct
type Integer struct {
	Value int64
}

// Inspect function
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type function
func (i *Integer) Type() Type { return INTEGER_OBJ }

// Boolean struct
type Boolean struct {
	Value bool
}

// Type func
func (b *Boolean) Type() Type { return BOOLEAN_OBJ }

// Inspect func
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null struct
type Null struct{}

// Type func
func (n *Null) Type() Type { return NULL_OBJ }

// Inspect func
func (n *Null) Inspect() string { return "null" }

// ReturnValue struct
type ReturnValue struct {
	Value Object
}

// Type func
func (rv *ReturnValue) Type() Type { return RETURN_VALUE_OBJ }

// Inspect func
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error struct
type Error struct {
	Message string
}

// Type func
func (e *Error) Type() Type { return ERROR_OBJ }

// Inspect func
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function struct
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type func
func (f *Function) Type() Type { return FUNCTION_OBJ }

// Inspect func
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
