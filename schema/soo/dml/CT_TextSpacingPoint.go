// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TextSpacingPoint struct {
	ValAttr int32
}

func NewCT_TextSpacingPoint() *CT_TextSpacingPoint {
	ret := &CT_TextSpacingPoint{}
	ret.ValAttr = 0
	return ret
}

func (m *CT_TextSpacingPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextSpacingPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = 0
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = int32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextSpacingPoint: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextSpacingPoint and its children
func (m *CT_TextSpacingPoint) Validate() error {
	return m.ValidateWithPath("CT_TextSpacingPoint")
}

// ValidateWithPath validates the CT_TextSpacingPoint and its children, prefixing error messages with path
func (m *CT_TextSpacingPoint) ValidateWithPath(path string) error {
	if m.ValAttr < 0 {
		return fmt.Errorf("%s/m.ValAttr must be >= 0 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr > 158400 {
		return fmt.Errorf("%s/m.ValAttr must be <= 158400 (have %v)", path, m.ValAttr)
	}
	return nil
}
