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

func TestCT_PathShadePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_PathShadeProperties()
	if v == nil {
		t.Errorf("dml.NewCT_PathShadeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PathShadeProperties should validate: %s", err)
	}
}

func TestCT_PathShadePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PathShadeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PathShadeProperties()
	xml.Unmarshal(buf, v2)
}
