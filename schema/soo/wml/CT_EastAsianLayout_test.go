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

func TestCT_EastAsianLayoutConstructor(t *testing.T) {
	v := wml.NewCT_EastAsianLayout()
	if v == nil {
		t.Errorf("wml.NewCT_EastAsianLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_EastAsianLayout should validate: %s", err)
	}
}

func TestCT_EastAsianLayoutMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_EastAsianLayout()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_EastAsianLayout()
	xml.Unmarshal(buf, v2)
}
