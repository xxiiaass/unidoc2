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

func TestCT_MetadataRecordConstructor(t *testing.T) {
	v := sml.NewCT_MetadataRecord()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataRecord must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataRecord should validate: %s", err)
	}
}

func TestCT_MetadataRecordMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataRecord()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataRecord()
	xml.Unmarshal(buf, v2)
}
