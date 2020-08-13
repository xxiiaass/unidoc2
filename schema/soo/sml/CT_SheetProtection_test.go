// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/sml"
)

func TestCT_SheetProtectionConstructor(t *testing.T) {
	v := sml.NewCT_SheetProtection()
	if v == nil {
		t.Errorf("sml.NewCT_SheetProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetProtection should validate: %s", err)
	}
}

func TestCT_SheetProtectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetProtection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetProtection()
	xml.Unmarshal(buf, v2)
}
