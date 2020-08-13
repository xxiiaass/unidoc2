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

	"github.com/xxiiaass/unidoc2/schema/soo/dml"
)

func TestCT_AudioFileConstructor(t *testing.T) {
	v := dml.NewCT_AudioFile()
	if v == nil {
		t.Errorf("dml.NewCT_AudioFile must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AudioFile should validate: %s", err)
	}
}

func TestCT_AudioFileMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AudioFile()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AudioFile()
	xml.Unmarshal(buf, v2)
}
