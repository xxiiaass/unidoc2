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

	"github.com/xxiiaass/unidoc2/schema/soo/wml"
)

func TestAG_PasswordConstructor(t *testing.T) {
	v := wml.NewAG_Password()
	if v == nil {
		t.Errorf("wml.NewAG_Password must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.AG_Password should validate: %s", err)
	}
}

func TestAG_PasswordMarshalUnmarshal(t *testing.T) {
	v := wml.NewAG_Password()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewAG_Password()
	xml.Unmarshal(buf, v2)
}
