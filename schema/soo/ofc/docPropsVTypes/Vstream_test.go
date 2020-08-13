// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/ofc/docPropsVTypes"
)

func TestVstreamConstructor(t *testing.T) {
	v := docPropsVTypes.NewVstream()
	if v == nil {
		t.Errorf("docPropsVTypes.NewVstream must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Vstream should validate: %s", err)
	}
}

func TestVstreamMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewVstream()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewVstream()
	xml.Unmarshal(buf, v2)
}