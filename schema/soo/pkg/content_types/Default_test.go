// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package content_types_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/pkg/content_types"
)

func TestDefaultConstructor(t *testing.T) {
	v := content_types.NewDefault()
	if v == nil {
		t.Errorf("content_types.NewDefault must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.Default should validate: %s", err)
	}
}

func TestDefaultMarshalUnmarshal(t *testing.T) {
	v := content_types.NewDefault()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewDefault()
	xml.Unmarshal(buf, v2)
}
