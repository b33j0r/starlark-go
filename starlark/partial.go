// partial.go
package starlark

import (
	"fmt"
)

// partialFunction represents a function with the first argument bound.
type partialFunction struct {
	original  Callable
	boundArgs []Value
}

// Ensure partialFunction implements Callable
var _ Callable = (*partialFunction)(nil)

// Type returns the type of the value.
func (p *partialFunction) Type() string {
	return "partial_function"
}

// Freeze makes the value immutable.
func (p *partialFunction) Freeze() {
	p.original.Freeze()
	for _, arg := range p.boundArgs {
		arg.Freeze()
	}
}

// Truth returns the truth value of the value.
func (p *partialFunction) Truth() Bool {
	return True
}

// Hash returns the hash of the value.
func (p *partialFunction) Hash() (uint32, error) {
	return 0, fmt.Errorf("partial_function: unhashable")
}

// String returns the string representation.
func (p *partialFunction) String() string {
	return fmt.Sprintf("(partial Fn=%s)", p.original.Name())
}

// Attr returns the attribute of the value.
func (p *partialFunction) Attr(name string) (Value, error) {
	return nil, nil
}

// AttrNames returns the list of attributes.
func (p *partialFunction) AttrNames() []string {
	return nil
}

// Name returns the name of the function.
func (p *partialFunction) Name() string {
	return p.original.Name() + "__partial"
}

// CallInternal invokes the partially applied function with remaining arguments.
func (p *partialFunction) CallInternal(thread *Thread, args Tuple, kwargs []Tuple) (Value, error) {
	// Prepend the bound argument to the provided arguments
	combinedArgs := make([]Value, len(p.boundArgs)+len(args))
	for i := range p.boundArgs {
		combinedArgs[i] = p.boundArgs[i]
	}
	for i := range args {
		combinedArgs[len(p.boundArgs)+i] = args[i]
	}

	// Debugging: Print the arguments being passed
	//fmt.Printf("Calling function '%s' with args: ", p.original.Name())
	//for _, arg := range combinedArgs {
	//	fmt.Printf("%s ", arg.String())
	//}
	//fmt.Println()

	return p.original.CallInternal(thread, Tuple(combinedArgs), kwargs)
}
