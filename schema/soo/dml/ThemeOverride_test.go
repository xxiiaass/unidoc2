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

	"github.com/unidoc/unioffice/schema/soo/dml"
)

func TestThemeOverrideConstructor(t *testing.T) {
	v := dml.NewThemeOverride()
	if v == nil {
		t.Errorf("dml.NewThemeOverride must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.ThemeOverride should validate: %s", err)
	}
}

func TestThemeOverrideMarshalUnmarshal(t *testing.T) {
	v := dml.NewThemeOverride()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewThemeOverride()
	xml.Unmarshal(buf, v2)
}