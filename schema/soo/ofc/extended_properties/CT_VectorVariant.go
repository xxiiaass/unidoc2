// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package extended_properties

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/ofc/docPropsVTypes"
)

type CT_VectorVariant struct {
	Vector *docPropsVTypes.Vector
}

func NewCT_VectorVariant() *CT_VectorVariant {
	ret := &CT_VectorVariant{}
	ret.Vector = docPropsVTypes.NewVector()
	return ret
}

func (m *CT_VectorVariant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sevector := xml.StartElement{Name: xml.Name{Local: "vt:vector"}}
	e.EncodeElement(m.Vector, sevector)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VectorVariant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Vector = docPropsVTypes.NewVector()
lCT_VectorVariant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "vector"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "vector"}:
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_VectorVariant %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VectorVariant
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VectorVariant and its children
func (m *CT_VectorVariant) Validate() error {
	return m.ValidateWithPath("CT_VectorVariant")
}

// ValidateWithPath validates the CT_VectorVariant and its children, prefixing error messages with path
func (m *CT_VectorVariant) ValidateWithPath(path string) error {
	if err := m.Vector.ValidateWithPath(path + "/Vector"); err != nil {
		return err
	}
	return nil
}