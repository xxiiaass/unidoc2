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

func TestCT_StyleDefinitionConstructor(t *testing.T) {
	v := diagram.NewCT_StyleDefinition()
	if v == nil {
		t.Errorf("diagram.NewCT_StyleDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_StyleDefinition should validate: %s", err)
	}
}

func TestCT_StyleDefinitionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_StyleDefinition()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_StyleDefinition()
	xml.Unmarshal(buf, v2)
}
