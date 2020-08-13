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

	"github.com/unidoc/unioffice/schema/soo/dml/chart"
)

func TestCT_RotYConstructor(t *testing.T) {
	v := chart.NewCT_RotY()
	if v == nil {
		t.Errorf("chart.NewCT_RotY must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RotY should validate: %s", err)
	}
}

func TestCT_RotYMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RotY()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RotY()
	xml.Unmarshal(buf, v2)
}