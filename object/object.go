// object/object.go

package object

import "fmt"

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
