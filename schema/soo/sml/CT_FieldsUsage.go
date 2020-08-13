// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/xxiiaass/unidoc2"
)

type CT_FieldsUsage struct {
	// Field Count
	CountAttr *uint32
	// PivotCache Field Id
	FieldUsage []*CT_FieldUsage
}

func NewCT_FieldsUsage() *CT_FieldsUsage {
	ret := &CT_FieldsUsage{}
	return ret
}

func (m *CT_FieldsUsage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.FieldUsage != nil {
		sefieldUsage := xml.StartElement{Name: xml.Name{Local: "ma:fieldUsage"}}
		for _, c := range m.FieldUsage {
			e.EncodeElement(c, sefieldUsage)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FieldsUsage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_FieldsUsage:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fieldUsage"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "fieldUsage"}:
				tmp := NewCT_FieldUsage()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FieldUsage = append(m.FieldUsage, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_FieldsUsage %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FieldsUsage
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FieldsUsage and its children
func (m *CT_FieldsUsage) Validate() error {
	return m.ValidateWithPath("CT_FieldsUsage")
}

// ValidateWithPath validates the CT_FieldsUsage and its children, prefixing error messages with path
func (m *CT_FieldsUsage) ValidateWithPath(path string) error {
	for i, v := range m.FieldUsage {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FieldUsage[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
