// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import (
	"sort"
	"sync"

	"github.com/xxiiaass/unidoc2"
)

// SupportedFunctions returns a list of supported functions.
func SupportedFunctions() []string {
	ret := []string{}
	for k := range registered {
		ret = append(ret, k)
	}
	for k := range registeredComplex {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

// Function is a standard function whose result only depends on its arguments.
type Function func(args []Result) Result

// FunctionComplex is a function whose result  depends on its arguments and the
// context that it's in.  As an example, INDIRECT is a complex function so that
// INDIRECT("A1") which returns the value of the "A1" cell in a sheet can use
// the context to reach into the sheet and pull out required values.
type FunctionComplex func(ctx Context, ev Evaluator, args []Result) Result

var regLock sync.Mutex
var registered = map[string]Function{}
var registeredComplex = map[string]FunctionComplex{}

// RegisterFunction registers a standard function.
func RegisterFunction(name string, fn Function) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, ok := registered[name]; ok {
		unioffice.Log("duplicate registration of function %s", name)
	}
	registered[name] = fn
}

// RegisterFunctionComplex registers a standard function.
func RegisterFunctionComplex(name string, fn FunctionComplex) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, ok := registeredComplex[name]; ok {
		unioffice.Log("duplicate registration of function %s", name)
	}
	registeredComplex[name] = fn
}

// LookupFunction looks up and returns a standard function or nil.
func LookupFunction(name string) Function {
	regLock.Lock()
	defer regLock.Unlock()
	if fn, ok := registered[name]; ok {
		return fn
	}
	return nil
}

// LookupFunctionComplex looks up and returns a complex function or nil.
func LookupFunctionComplex(name string) FunctionComplex {
	regLock.Lock()
	defer regLock.Unlock()
	if fn, ok := registeredComplex[name]; ok {
		return fn
	}
	return nil
}
