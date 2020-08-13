// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import "github.com/xxiiaass/unidoc2/spreadsheet/update"

// ConstArrayExpr is a constant array expression.
type ConstArrayExpr struct {
	data [][]Expression
}

// NewConstArrayExpr constructs a new constant array expression with a given data.
func NewConstArrayExpr(data [][]Expression) Expression {
	return &ConstArrayExpr{data}
}

// Eval evaluates and returns the result of a constant array expression.
func (c ConstArrayExpr) Eval(ctx Context, ev Evaluator) Result {
	res := [][]Result{}
	for _, row := range c.data {
		r := []Result{}
		for _, col := range row {
			r = append(r, col.Eval(ctx, ev))
		}
		res = append(res, r)
	}
	return MakeArrayResult(res)
}

// Reference returns an invalid reference for ConstArrayExpr.
func (c ConstArrayExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// String returns a string representation of ConstArrayExpr.
func (c ConstArrayExpr) String() string {
	return "" // to do
}

// Update returns the same object as updating sheet references does not affect ConstArrayExpr.
func (c ConstArrayExpr) Update(q *update.UpdateQuery) Expression {
	return c
}
