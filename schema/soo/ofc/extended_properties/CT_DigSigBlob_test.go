// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/ofc/extended_properties"
)

func TestCT_DigSigBlobConstructor(t *testing.T) {
	v := extended_properties.NewCT_DigSigBlob()
	if v == nil {
		t.Errorf("extended_properties.NewCT_DigSigBlob must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.CT_DigSigBlob should validate: %s", err)
	}
}

func TestCT_DigSigBlobMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewCT_DigSigBlob()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewCT_DigSigBlob()
	xml.Unmarshal(buf, v2)
}
