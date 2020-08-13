// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/ofc/math"
)

func TestCT_ManualBreakConstructor(t *testing.T) {
	v := math.NewCT_ManualBreak()
	if v == nil {
		t.Errorf("math.NewCT_ManualBreak must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_ManualBreak should validate: %s", err)
	}
}

func TestCT_ManualBreakMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_ManualBreak()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_ManualBreak()
	xml.Unmarshal(buf, v2)
}
