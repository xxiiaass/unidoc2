// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/ofc/custom_properties"
)

func TestCT_PropertiesConstructor(t *testing.T) {
	v := custom_properties.NewCT_Properties()
	if v == nil {
		t.Errorf("custom_properties.NewCT_Properties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.CT_Properties should validate: %s", err)
	}
}

func TestCT_PropertiesMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewCT_Properties()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewCT_Properties()
	xml.Unmarshal(buf, v2)
}
