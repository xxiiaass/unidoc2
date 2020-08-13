// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/dml"
)

func TestCT_TextFieldConstructor(t *testing.T) {
	v := dml.NewCT_TextField()
	if v == nil {
		t.Errorf("dml.NewCT_TextField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextField should validate: %s", err)
	}
}

func TestCT_TextFieldMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextField()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextField()
	xml.Unmarshal(buf, v2)
}
