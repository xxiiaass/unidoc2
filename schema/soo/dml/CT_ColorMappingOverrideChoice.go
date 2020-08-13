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

	"github.com/unidoc/unioffice"
)

type CT_ColorMappingOverrideChoice struct {
	MasterClrMapping   *CT_EmptyElement
	OverrideClrMapping *CT_ColorMapping
}

func NewCT_ColorMappingOverrideChoice() *CT_ColorMappingOverrideChoice {
	ret := &CT_ColorMappingOverrideChoice{}
	return ret
}

func (m *CT_ColorMappingOverrideChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MasterClrMapping != nil {
		semasterClrMapping := xml.StartElement{Name: xml.Name{Local: "a:masterClrMapping"}}
		e.EncodeElement(m.MasterClrMapping, semasterClrMapping)
	}
	if m.OverrideClrMapping != nil {
		seoverrideClrMapping := xml.StartElement{Name: xml.Name{Local: "a:overrideClrMapping"}}
		e.EncodeElement(m.OverrideClrMapping, seoverrideClrMapping)
	}
	return nil
}

func (m *CT_ColorMappingOverrideChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorMappingOverrideChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "masterClrMapping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "masterClrMapping"}:
				m.MasterClrMapping = NewCT_EmptyElement()
				if err := d.DecodeElement(m.MasterClrMapping, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "overrideClrMapping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "overrideClrMapping"}:
				m.OverrideClrMapping = NewCT_ColorMapping()
				if err := d.DecodeElement(m.OverrideClrMapping, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ColorMappingOverrideChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorMappingOverrideChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorMappingOverrideChoice and its children
func (m *CT_ColorMappingOverrideChoice) Validate() error {
	return m.ValidateWithPath("CT_ColorMappingOverrideChoice")
}

// ValidateWithPath validates the CT_ColorMappingOverrideChoice and its children, prefixing error messages with path
func (m *CT_ColorMappingOverrideChoice) ValidateWithPath(path string) error {
	if m.MasterClrMapping != nil {
		if err := m.MasterClrMapping.ValidateWithPath(path + "/MasterClrMapping"); err != nil {
			return err
		}
	}
	if m.OverrideClrMapping != nil {
		if err := m.OverrideClrMapping.ValidateWithPath(path + "/OverrideClrMapping"); err != nil {
			return err
		}
	}
	return nil
}