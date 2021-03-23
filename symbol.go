package gslang

// SymbolScope represents a symbol scope.
type SymbolScope string

// List of symbol scopes
const (
	ScopeGlobal  SymbolScope = "GLOBAL"
	ScopeLocal   SymbolScope = "LOCAL"
	ScopeBuiltin SymbolScope = "BUILTIN"
	ScopeFree    SymbolScope = "FREE"
)

// SymbolObject represents a symbol in the symbol table.
type SymbolObject struct {
	Name          string
	Scope         SymbolScope
	Index         int
	LocalAssigned bool // if the local symbol is assigned at least once
}

// Symbol represents a symbol table.
type Symbol struct {
	parent         *Symbol
	block          bool
	store          map[string]*SymbolObject
	numDefinition  int
	maxDefinition  int
	freeSymbols    []*SymbolObject
	builtinSymbols []*SymbolObject
}

// NewSymbol creates a Symbol.
func NewSymbol() *Symbol {
	return &Symbol{
		store: make(map[string]*SymbolObject),
	}
}

// Define adds a new symbol in the current scope.
func (t *Symbol) Define(name string) *SymbolObject {
	symbol := &SymbolObject{Name: name, Index: t.nextIndex()}
	t.numDefinition++

	if t.Parent(true) == nil {
		symbol.Scope = ScopeGlobal

		// if symbol is defined in a block of global scope, symbol index must
		// be tracked at the root-level table instead.
		if p := t.parent; p != nil {
			for p.parent != nil {
				p = p.parent
			}
			t.numDefinition--
			p.numDefinition++
		}

	} else {
		symbol.Scope = ScopeLocal
	}
	t.store[name] = symbol
	t.updateMaxDefs(symbol.Index + 1)
	return symbol
}

// DefineBuiltin adds a symbol for builtin function.
func (t *Symbol) DefineBuiltin(index int, name string) *SymbolObject {
	if t.parent != nil {
		return t.parent.DefineBuiltin(index, name)
	}

	symbol := &SymbolObject{
		Name:  name,
		Index: index,
		Scope: ScopeBuiltin,
	}
	t.store[name] = symbol
	t.builtinSymbols = append(t.builtinSymbols, symbol)
	return symbol
}

// Resolve resolves a symbol with a given name.
func (t *Symbol) Resolve(
	name string,
	recur bool,
) (*SymbolObject, int, bool) {
	symbol, ok := t.store[name]
	if ok {
		// symbol can be used if
		if symbol.Scope != ScopeLocal || // it's not of local scope, OR,
			symbol.LocalAssigned || // it's assigned at least once, OR,
			recur { // it's defined in higher level
			return symbol, 0, true
		}
	}

	if t.parent == nil {
		return nil, 0, false
	}

	symbol, depth, ok := t.parent.Resolve(name, true)
	if !ok {
		return nil, 0, false
	}
	depth++

	// if symbol is defined in parent table and if it's not global/builtin
	// then it's free variable.
	if !t.block && depth > 0 &&
		symbol.Scope != ScopeGlobal &&
		symbol.Scope != ScopeBuiltin {
		return t.defineFree(symbol), depth, true
	}
	return symbol, depth, true
}

// Fork creates a new symbol table for a new scope.
func (t *Symbol) Fork(block bool) *Symbol {
	return &Symbol{
		store:  make(map[string]*SymbolObject),
		parent: t,
		block:  block,
	}
}

// Parent returns the outer scope of the current symbol table.
func (t *Symbol) Parent(skipBlock bool) *Symbol {
	if skipBlock && t.block {
		return t.parent.Parent(skipBlock)
	}
	return t.parent
}

// MaxSymbols returns the total number of symbols defined in the scope.
func (t *Symbol) MaxSymbols() int {
	return t.maxDefinition
}

// FreeSymbols returns free symbols for the scope.
func (t *Symbol) FreeSymbols() []*SymbolObject {
	return t.freeSymbols
}

// BuiltinSymbols returns builtin symbols for the scope.
func (t *Symbol) BuiltinSymbols() []*SymbolObject {
	if t.parent != nil {
		return t.parent.BuiltinSymbols()
	}
	return t.builtinSymbols
}

// Names returns the name of all the symbols.
func (t *Symbol) Names() []string {
	var names []string
	for name := range t.store {
		names = append(names, name)
	}
	return names
}

func (t *Symbol) nextIndex() int {
	if t.block {
		return t.parent.nextIndex() + t.numDefinition
	}
	return t.numDefinition
}

func (t *Symbol) updateMaxDefs(numDefs int) {
	if numDefs > t.maxDefinition {
		t.maxDefinition = numDefs
	}
	if t.block {
		t.parent.updateMaxDefs(numDefs)
	}
}

func (t *Symbol) defineFree(original *SymbolObject) *SymbolObject {
	// TODO: should we check duplicates?
	t.freeSymbols = append(t.freeSymbols, original)
	symbol := &SymbolObject{
		Name:  original.Name,
		Index: len(t.freeSymbols) - 1,
		Scope: ScopeFree,
	}
	t.store[original.Name] = symbol
	return symbol
}
