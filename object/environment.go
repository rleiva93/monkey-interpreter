package object

// NewEnvironment func
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment struct
type Environment struct {
	store map[string]Object
}

// Get func
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set func
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
