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

func TestEG_HdrFtrReferencesConstructor(t *testing.T) {
	v := wml.NewEG_HdrFtrReferences()
	if v == nil {
		t.Errorf("wml.NewEG_HdrFtrReferences must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_HdrFtrReferences should validate: %s", err)
	}
}

func TestEG_HdrFtrReferencesMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_HdrFtrReferences()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_HdrFtrReferences()
	xml.Unmarshal(buf, v2)
}
