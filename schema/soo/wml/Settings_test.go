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

func TestSettingsConstructor(t *testing.T) {
	v := wml.NewSettings()
	if v == nil {
		t.Errorf("wml.NewSettings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Settings should validate: %s", err)
	}
}

func TestSettingsMarshalUnmarshal(t *testing.T) {
	v := wml.NewSettings()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewSettings()
	xml.Unmarshal(buf, v2)
}