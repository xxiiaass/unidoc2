// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package word_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/urn/schemas_microsoft_com/office/word"
)

func TestBorderleftConstructor(t *testing.T) {
	v := word.NewBorderleft()
	if v == nil {
		t.Errorf("word.NewBorderleft must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Borderleft should validate: %s", err)
	}
}

func TestBorderleftMarshalUnmarshal(t *testing.T) {
	v := word.NewBorderleft()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBorderleft()
	xml.Unmarshal(buf, v2)
}
