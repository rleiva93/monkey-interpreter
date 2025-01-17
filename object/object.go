// object/object.go

package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
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
	// STRING_OBJ const
	STRING_OBJ = "STRING"
	// BUILTIN_OBJ const
	BUILTIN_OBJ = "BUILTIN"
	// ARRAY_OBJ const
	ARRAY_OBJ = "ARRAY"
	// HASH_OBJ const
	HASH_OBJ = "HASH"
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

// String struct
type String struct {
	Value string
}

// Type func
func (s *String) Type() Type { return STRING_OBJ }

// Inspect func
func (s *String) Inspect() string { return s.Value }

// BuiltinFunction type
type BuiltinFunction func(args ...Object) Object

// Builtin struct
type Builtin struct {
	Fn BuiltinFunction
}

// Type func
func (b *Builtin) Type() Type { return BUILTIN_OBJ }

// Inspect func
func (b *Builtin) Inspect() string { return "builtin function" }

// Array struct
type Array struct {
	Elements []Object
}

// Type func
func (ao *Array) Type() Type { return ARRAY_OBJ }

// Inspect func
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey struct
type HashKey struct {
	Type  Type
	Value uint64
}

// HashKey func
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey func
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey func
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair struct
type HashPair struct {
	Key   Object
	Value Object
}

// Hash struct
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type func
func (h *Hash) Type() Type { return HASH_OBJ }

// Inspect func
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Hashable interface
type Hashable interface {
	HashKey() HashKey
}
