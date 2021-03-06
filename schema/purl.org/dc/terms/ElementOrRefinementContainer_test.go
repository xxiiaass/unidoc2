// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/xxiiaass/unidoc2/schema/purl.org/dc/terms"
)

func TestElementOrRefinementContainerConstructor(t *testing.T) {
	v := terms.NewElementOrRefinementContainer()
	if v == nil {
		t.Errorf("terms.NewElementOrRefinementContainer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ElementOrRefinementContainer should validate: %s", err)
	}
}

func TestElementOrRefinementContainerMarshalUnmarshal(t *testing.T) {
	v := terms.NewElementOrRefinementContainer()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewElementOrRefinementContainer()
	xml.Unmarshal(buf, v2)
}
