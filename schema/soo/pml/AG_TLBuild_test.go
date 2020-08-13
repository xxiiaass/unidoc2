// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/pml"
)

func TestAG_TLBuildConstructor(t *testing.T) {
	v := pml.NewAG_TLBuild()
	if v == nil {
		t.Errorf("pml.NewAG_TLBuild must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.AG_TLBuild should validate: %s", err)
	}
}

func TestAG_TLBuildMarshalUnmarshal(t *testing.T) {
	v := pml.NewAG_TLBuild()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewAG_TLBuild()
	xml.Unmarshal(buf, v2)
}
