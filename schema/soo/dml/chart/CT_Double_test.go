// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/dml/chart"
)

func TestCT_DoubleConstructor(t *testing.T) {
	v := chart.NewCT_Double()
	if v == nil {
		t.Errorf("chart.NewCT_Double must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Double should validate: %s", err)
	}
}

func TestCT_DoubleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Double()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Double()
	xml.Unmarshal(buf, v2)
}
