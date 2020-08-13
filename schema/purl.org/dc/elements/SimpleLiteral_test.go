// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package elements_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/purl.org/dc/elements"
)

func TestSimpleLiteralConstructor(t *testing.T) {
	v := elements.NewSimpleLiteral()
	if v == nil {
		t.Errorf("elements.NewSimpleLiteral must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.SimpleLiteral should validate: %s", err)
	}
}

func TestSimpleLiteralMarshalUnmarshal(t *testing.T) {
	v := elements.NewSimpleLiteral()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewSimpleLiteral()
	xml.Unmarshal(buf, v2)
}
