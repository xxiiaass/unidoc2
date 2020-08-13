// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/dml/chartDrawing"
)

func TestCT_ShapeNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_ShapeNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_ShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_ShapeNonVisual should validate: %s", err)
	}
}

func TestCT_ShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_ShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_ShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
