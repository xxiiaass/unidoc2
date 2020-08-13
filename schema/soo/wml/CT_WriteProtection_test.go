// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/wml"
)

func TestCT_WriteProtectionConstructor(t *testing.T) {
	v := wml.NewCT_WriteProtection()
	if v == nil {
		t.Errorf("wml.NewCT_WriteProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_WriteProtection should validate: %s", err)
	}
}

func TestCT_WriteProtectionMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_WriteProtection()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_WriteProtection()
	xml.Unmarshal(buf, v2)
}
