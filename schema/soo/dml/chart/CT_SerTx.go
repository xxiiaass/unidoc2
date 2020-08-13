// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_SerTx struct {
	Choice *CT_SerTxChoice
}

func NewCT_SerTx() *CT_SerTx {
	ret := &CT_SerTx{}
	ret.Choice = NewCT_SerTxChoice()
	return ret
}

func (m *CT_SerTx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, xml.StartElement{})
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SerTx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_SerTxChoice()
lCT_SerTx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strRef"}:
				m.Choice = NewCT_SerTxChoice()
				if err := d.DecodeElement(&m.Choice.StrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "v"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "v"}:
				m.Choice = NewCT_SerTxChoice()
				if err := d.DecodeElement(&m.Choice.V, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SerTx %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SerTx
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SerTx and its children
func (m *CT_SerTx) Validate() error {
	return m.ValidateWithPath("CT_SerTx")
}

// ValidateWithPath validates the CT_SerTx and its children, prefixing error messages with path
func (m *CT_SerTx) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
