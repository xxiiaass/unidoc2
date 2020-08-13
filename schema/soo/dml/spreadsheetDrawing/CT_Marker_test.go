// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_MarkerConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Marker()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Marker must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Marker should validate: %s", err)
	}
}

func TestCT_MarkerMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Marker()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Marker()
	xml.Unmarshal(buf, v2)
}
