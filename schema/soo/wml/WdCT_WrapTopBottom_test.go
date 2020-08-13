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

	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestWdCT_WrapTopBottomConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapTopBottom()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapTopBottom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapTopBottom should validate: %s", err)
	}
}

func TestWdCT_WrapTopBottomMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapTopBottom()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapTopBottom()
	xml.Unmarshal(buf, v2)
}