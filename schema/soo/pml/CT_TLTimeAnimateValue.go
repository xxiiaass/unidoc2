// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/xxiiaass/unidoc2"
)

type CT_TLTimeAnimateValue struct {
	// Time
	TmAttr *ST_TLTimeAnimateValueTime
	// Formula
	FmlaAttr *string
	// Value
	Val *CT_TLAnimVariant
}

func NewCT_TLTimeAnimateValue() *CT_TLTimeAnimateValue {
	ret := &CT_TLTimeAnimateValue{}
	return ret
}

func (m *CT_TLTimeAnimateValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tm"},
			Value: fmt.Sprintf("%v", *m.TmAttr)})
	}
	if m.FmlaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fmla"},
			Value: fmt.Sprintf("%v", *m.FmlaAttr)})
	}
	e.EncodeToken(start)
	if m.Val != nil {
		seval := xml.StartElement{Name: xml.Name{Local: "p:val"}}
		e.EncodeElement(m.Val, seval)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeAnimateValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tm" {
			parsed, err := ParseUnionST_TLTimeAnimateValueTime(attr.Value)
			if err != nil {
				return err
			}
			m.TmAttr = &parsed
			continue
		}
		if attr.Name.Local == "fmla" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FmlaAttr = &parsed
			continue
		}
	}
lCT_TLTimeAnimateValue:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "val"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "val"}:
				m.Val = NewCT_TLAnimVariant()
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TLTimeAnimateValue %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeAnimateValue
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeAnimateValue and its children
func (m *CT_TLTimeAnimateValue) Validate() error {
	return m.ValidateWithPath("CT_TLTimeAnimateValue")
}

// ValidateWithPath validates the CT_TLTimeAnimateValue and its children, prefixing error messages with path
func (m *CT_TLTimeAnimateValue) ValidateWithPath(path string) error {
	if m.TmAttr != nil {
		if err := m.TmAttr.ValidateWithPath(path + "/TmAttr"); err != nil {
			return err
		}
	}
	if m.Val != nil {
		if err := m.Val.ValidateWithPath(path + "/Val"); err != nil {
			return err
		}
	}
	return nil
}
