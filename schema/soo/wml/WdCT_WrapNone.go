// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"
)

type WdCT_WrapNone struct {
}

func NewWdCT_WrapNone() *WdCT_WrapNone {
	ret := &WdCT_WrapNone{}
	return ret
}

func (m *WdCT_WrapNone) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WrapNone) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing WdCT_WrapNone: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the WdCT_WrapNone and its children
func (m *WdCT_WrapNone) Validate() error {
	return m.ValidateWithPath("WdCT_WrapNone")
}

// ValidateWithPath validates the WdCT_WrapNone and its children, prefixing error messages with path
func (m *WdCT_WrapNone) ValidateWithPath(path string) error {
	return nil
}