// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/pml"
)

func TestCT_TLShapeTargetElementConstructor(t *testing.T) {
	v := pml.NewCT_TLShapeTargetElement()
	if v == nil {
		t.Errorf("pml.NewCT_TLShapeTargetElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLShapeTargetElement should validate: %s", err)
	}
}

func TestCT_TLShapeTargetElementMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLShapeTargetElement()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLShapeTargetElement()
	xml.Unmarshal(buf, v2)
}
