// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_RoundRectConstructor(t *testing.T) {
	v := vml.NewCT_RoundRect()
	if v == nil {
		t.Errorf("vml.NewCT_RoundRect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_RoundRect should validate: %s", err)
	}
}

func TestCT_RoundRectMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_RoundRect()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_RoundRect()
	xml.Unmarshal(buf, v2)
}
