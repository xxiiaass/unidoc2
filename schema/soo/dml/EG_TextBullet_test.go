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

func TestEG_TextBulletConstructor(t *testing.T) {
	v := dml.NewEG_TextBullet()
	if v == nil {
		t.Errorf("dml.NewEG_TextBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextBullet should validate: %s", err)
	}
}

func TestEG_TextBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextBullet()
	xml.Unmarshal(buf, v2)
}
