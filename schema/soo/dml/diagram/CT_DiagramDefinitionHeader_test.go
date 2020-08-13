// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/soo/dml/diagram"
)

func TestCT_DiagramDefinitionHeaderConstructor(t *testing.T) {
	v := diagram.NewCT_DiagramDefinitionHeader()
	if v == nil {
		t.Errorf("diagram.NewCT_DiagramDefinitionHeader must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_DiagramDefinitionHeader should validate: %s", err)
	}
}

func TestCT_DiagramDefinitionHeaderMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_DiagramDefinitionHeader()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_DiagramDefinitionHeader()
	xml.Unmarshal(buf, v2)
}
