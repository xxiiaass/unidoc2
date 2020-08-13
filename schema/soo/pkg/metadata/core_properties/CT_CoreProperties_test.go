// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package core_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/pkg/metadata/core_properties"
)

func TestCT_CorePropertiesConstructor(t *testing.T) {
	v := core_properties.NewCT_CoreProperties()
	if v == nil {
		t.Errorf("core_properties.NewCT_CoreProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed core_properties.CT_CoreProperties should validate: %s", err)
	}
}

func TestCT_CorePropertiesMarshalUnmarshal(t *testing.T) {
	v := core_properties.NewCT_CoreProperties()
	buf, _ := xml.Marshal(v)
	v2 := core_properties.NewCT_CoreProperties()
	xml.Unmarshal(buf, v2)
}
