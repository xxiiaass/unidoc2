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

	"github.com/unidoc/unioffice/schema/soo/dml/chartDrawing"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Shape()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}