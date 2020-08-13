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

	"github.com/unidoc/unioffice"
)

type CT_OleObjects struct {
	// Embedded Object
	OleObject []*CT_OleObject
}

func NewCT_OleObjects() *CT_OleObjects {
	ret := &CT_OleObjects{}
	return ret
}

func (m *CT_OleObjects) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seoleObject := xml.StartElement{Name: xml.Name{Local: "ma:oleObject"}}
	for _, c := range m.OleObject {
		e.EncodeElement(c, seoleObject)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObjects) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OleObjects:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleObject"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "oleObject"}:
				tmp := NewCT_OleObject()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.OleObject = append(m.OleObject, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_OleObjects %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObjects
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObjects and its children
func (m *CT_OleObjects) Validate() error {
	return m.ValidateWithPath("CT_OleObjects")
}

// ValidateWithPath validates the CT_OleObjects and its children, prefixing error messages with path
func (m *CT_OleObjects) ValidateWithPath(path string) error {
	for i, v := range m.OleObject {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/OleObject[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}